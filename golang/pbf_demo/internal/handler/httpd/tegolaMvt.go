package httpd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/mvt"
	"github.com/go-spatial/geom/encoding/wkb"
	"github.com/go-spatial/geom/slippy"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"log"
	"net/http"
	"pbf_demo/internal/logger"
	"pbf_demo/internal/model"
	"pbf_demo/internal/params"
	"pbf_demo/internal/utils/response"
	"pbf_demo/internal/utils/tegola/convert"
	"pbf_demo/internal/utils/tegola/maths/validate"
	"strconv"
	"sync"
	"time"
)

var clipRegion = geom.NewExtent([2]float64{-64, -64}, [2]float64{4160, 4160})

// @Summary tegola
// @Description  tegola mvt
// @version 1.0
// @tags mvt
// @Param z path int true "z"
// @Param x path int true "x"
// @Param y path int true "y"
// @Success 200 {object} []byte "{"code":200,"data": "","msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"bind query err/params error"}"
// @Router /tegola/{z}/{x}/{y} [get]
func TegolaMVTGet(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	q := params.MVTGet{}
	if err := c.BindUri(&q); err != nil {
		logger.Logger.Error("bind uri err",
			zap.Error(err))
		utilGin.Response(400, "bind uri err", nil)
		return
	}
	x, _ := strconv.ParseUint(q.X, 10, 32)
	y, _ := strconv.ParseUint(q.Y, 10, 32)
	z, _ := strconv.ParseUint(q.Z, 10, 32)
	xyz := slippy.NewTile(uint(z), uint(x), uint(y))
	extent := xyz.Extent4326()
	e3857 := xyz.Extent3857()
	sql := fmt.Sprintf(`SELECT ST_AsBinary(st_transform(the_geom, 3857)) AS the_geom FROM adcode_test
WHERE the_geom && ST_MakeEnvelope(%v,%v,%v,%v, 4326)`, extent[0], extent[1], extent[2], extent[3])
	rows, err := model.DB.Debug().Raw(sql).Rows()
	if err != nil {
		logger.Logger.Error("DB error",
			zap.Error(err))
		utilGin.Response(400, "DB error", nil)
		return
	}
	var wg sync.WaitGroup
	var t mvt.Tile
	layer := mvt.Layer{
		Name: "adcode",
	}
	gs := make([]geom.Geometry, 0)
	ctx := context.Background()
	var geobytes []byte
	for rows.Next() {
		rows.Scan(&geobytes)
		geo, err := wkb.DecodeBytes(geobytes)
		if err != nil {
			logger.Logger.Error("wkb DecodeBytes error",
				zap.Error(err))
			utilGin.Response(400, "wkb DecodeBytes error", nil)
			return
		}
		gs = append(gs, geo)
	}
	start := time.Now()
	fmt.Printf("开始并行计算\n")
	for _, geo := range gs {
		wg.Add(1)
		go func(geo geom.Geometry, wg *sync.WaitGroup) {
			geo = mvt.PrepareGeo(geo, e3857, float64(mvt.DefaultExtent))
			defer wg.Done()
			sg, err := convert.ToTegola(geo)
			if err != nil {
				log.Fatalln("convert tegola error, err: ", err)
			}
			tegolaGeo, err := validate.CleanGeometry(ctx, sg, clipRegion)
			if err != nil {
				log.Fatalln("clean geometry error, err: ", err)
			}
			geo, err = convert.ToGeom(tegolaGeo)
			if err != nil {
				log.Fatalln("convert geometry error, err: ", err)
			}
			layer.AddFeatures(mvt.Feature{Geometry: geo})
		}(geo, &wg)
	}
	wg.Wait()
	fmt.Printf("并行计算结束, 瓦片编号: %v 耗时: %v\n", q, time.Since(start).Seconds())
	if err = t.AddLayers(&layer); err != nil {
		logger.Logger.Error("add layer error",
			zap.Error(err))
		utilGin.Response(400, "add layer error", nil)
		return
	}
	vtile, err := t.VTile(context.Background())
	if err != nil {
		logger.Logger.Error("t VTile error",
			zap.Error(err))
		utilGin.Response(400, "t VTile error", nil)
		return
	}
	pbf := make([]byte, 0)
	pbf, err = proto.Marshal(vtile)
	if err != nil {
		logger.Logger.Error("proto Marshal error",
			zap.Error(err))
		utilGin.Response(400, "proto Marshal error", nil)
		return
	}
	//fmt.Println(p)
	//	//// buffer to store our compressed bytes
	//	//var gzipBuf bytes.Buffer
	//	//
	//	//// compress the encoded bytes
	//	//w := gzip.NewWriter(&gzipBuf)
	//	//_, err = w.Write(pbf)
	//	//if err != nil {
	//	//	logger.Logger.Error("Write error",
	//	//		zap.Error(err))
	//	//	utilGin.Response(400, "Write error", nil)
	//	//	return
	//	//}
	//	//
	//	//// flush and close the writer
	//	//if err = w.Close(); err != nil {
	//	//	logger.Logger.Error("Close error",
	//	//		zap.Error(err))
	//	//	utilGin.Response(400, "Close error", nil)
	//	//	return
	//	//}
	//	//
	//	//pbf = gzipBuf.Bytes()
	//	//c.Header("Content-Encoding", "gzip")
	c.Header("Content-Length", fmt.Sprintf("%d", len(pbf)))
	c.Data(http.StatusOK, "application/vnd.mapbox-vector-tile", pbf)
	return
}
