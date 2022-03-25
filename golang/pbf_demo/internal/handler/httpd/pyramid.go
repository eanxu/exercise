package httpd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
	"strconv"
)

// @Summary pyramid
// @Description  pyramid
// @version 1.0
// @tags mvt
// @Param z path string true "z"
// @Param x path string true "x"
// @Param y path string true "y"
// @Success 200 {object} []byte "{"code":200,"data": "","msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"bind query err/params error"}"
// @Router /pyramid/{z}/{x}/{y} [get]
func PyramidMVTGet(c *gin.Context) {
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
	tile := slippy.NewTile(uint(z), uint(x), uint(y))
	var j string
	pbf := make([]byte, 0)
	model.DB.Debug().Raw(`select pyramid from adcode_test_pyramid where x=? and y=? and z=?`, tile.X, tile.Y, tile.Z).Scan(&j)
	if len(j) == 0 {
		c.Header("Content-Length", fmt.Sprintf("%d", len(pbf)))
		c.Data(http.StatusOK, "application/vnd.mapbox-vector-tile", pbf)
		return
	}
	newJ := make(map[uint64][]byte, 0)
	err := json.Unmarshal([]byte(j), &newJ)
	if err != nil {
		log.Fatalln("json unmarshal error, err: ", err)
	}
	ids := make([]uint64, 0)
	for k, _ := range newJ {
		ids = append(ids, k)
	}
	rows, err := model.DB.Raw(`select fid, "总就业", "第一产", "第二产" from adcode_test where fid in (?)`, ids).Rows()
	if err != nil {
		log.Fatalln("search data error, err: ", err)
	}
	var (
		a, b, d string
		fid uint64
	)
	var t mvt.Tile
		layer := mvt.Layer{
			Name: "adcode",
		}
	for rows.Next() {
		rows.Scan(&fid, &a, &b, &d)
		g, err := wkb.DecodeBytes(newJ[fid])
		if err != nil {
			log.Fatalln("wkb decode error, err: ", err)
		}
		feature := mvt.Feature{
			ID:       &fid,
			Tags:     map[string]interface{}{
				"总就业": a, "第一产":b, "第二产":d,
			},
			Geometry: g,
		}
		layer.AddFeatures(feature)
	}

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
	pbf, err = proto.Marshal(vtile)
	if err != nil {
		logger.Logger.Error("proto Marshal error",
			zap.Error(err))
		utilGin.Response(400, "proto Marshal error", nil)
		return
	}
	//// buffer to store our compressed bytes
	//var gzipBuf bytes.Buffer
	//
	//// compress the encoded bytes
	//w := gzip.NewWriter(&gzipBuf)
	//_, err = w.Write(pbf)
	//if err != nil {
	//	logger.Logger.Error("Write error",
	//		zap.Error(err))
	//	utilGin.Response(400, "Write error", nil)
	//	return
	//}
	//
	//// flush and close the writer
	//if err = w.Close(); err != nil {
	//	logger.Logger.Error("Close error",
	//		zap.Error(err))
	//	utilGin.Response(400, "Close error", nil)
	//	return
	//}
	//
	//pbf = gzipBuf.Bytes()
	//c.Header("Content-Encoding", "gzip")
	c.Header("Content-Length", fmt.Sprintf("%d", len(pbf)))
	c.Data(http.StatusOK, "application/vnd.mapbox-vector-tile", pbf)
	return
}
