package httpd

import (
	"fmt"
	"github.com/dadadamarine/orb"
	"github.com/dadadamarine/orb/encoding/mvt"
	"github.com/dadadamarine/orb/encoding/wkb"
	"github.com/dadadamarine/orb/geojson"
	"github.com/dadadamarine/orb/maptile"
	"github.com/dadadamarine/orb/simplify"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"pbf_demo/internal/logger"
	"pbf_demo/internal/model"
	"pbf_demo/internal/params"
	"pbf_demo/internal/utils/response"
	"strconv"
)

// @Summary orb
// @Description  orb包 mvt demo
// @version 1.0
// @tags mvt
// @Param z path int true "z"
// @Param x path int true "x"
// @Param y path int true "y"
// @Success 200 {object} []byte "{"code":200,"data": "","msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"bind query err/params error"}"
// @Router /orb/{z}/{x}/{y} [get]
func OrbMvtGet(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	q := params.MVTGet{}
	if err := c.BindUri(&q); err != nil {
		logger.Logger.Error("bind uri err",
			zap.Error(err))
		utilGin.Response(400, "bind uri err", nil)
		return
	}
	pbf := make([]byte, 0)
	x, _ := strconv.ParseUint(q.X, 10, 32)
	y, _ := strconv.ParseUint(q.Y, 10, 32)
	z, _ := strconv.ParseUint(q.Z, 10, 32)
	tile := maptile.New(uint32(x), uint32(y), maptile.Zoom(z))
	b := tile.Bound()
	//sql := fmt.Sprintf(`SELECT fid, ST_AsBinary(st_makevalid(the_geom)) AS the_geom FROM adcode_test WHERE the_geom && ST_MakeEnvelope(%v,%v,%v,%v, 4326)`, b.Left(), b.Top(), b.Right(), b.Bottom())
	sql := fmt.Sprintf(`SELECT fid, ST_AsBinary(st_makevalid(the_geom)) AS the_geom FROM adcode_test WHERE the_geom && ST_MakeEnvelope(%v,%v,%v,%v, 4326)`, b.Min[0], b.Min[1], b.Max[0], b.Max[1])
	rows, err := model.DB.Debug().Raw(sql).Rows()
	if err != nil {
		logger.Logger.Error("DB error",
			zap.Error(err))
		utilGin.Response(400, "DB error", nil)
		return
	}
	var (
		fid int
		geobytes []byte
		g orb.Geometry
	)
	fc := geojson.NewFeatureCollection()
	for rows.Next() {
		rows.Scan(&fid, &geobytes)
		g, err = wkb.Unmarshal(geobytes)
		if err != nil {
			logger.Logger.Error("wkb Unmarshal error",
				zap.Error(err))
			utilGin.Response(400, "wkb Unmarshal error", nil)
			return
		}
		feature := geojson.NewFeature(g)
		feature.ID = fid
		fc.Append(feature)
	}
	collections := map[string]*geojson.FeatureCollection{"adcode": fc}
	layers := mvt.NewLayers(collections)
	layers.ProjectToTile(tile)

	layers.Simplify(simplify.DouglasPeucker(1.0))
	// encoding using the Mapbox Vector Tile protobuf encoding.
	pbf, err = mvt.Marshal(layers) // this data is NOT gzipped.

	// error checking
	if err != nil {
		log.Fatalf("marshal error: %v", err)
	}
	c.Header("Content-Length", fmt.Sprintf("%d", len(pbf)))
	c.Data(http.StatusOK, "application/vnd.mapbox-vector-tile", pbf)
	return
}

