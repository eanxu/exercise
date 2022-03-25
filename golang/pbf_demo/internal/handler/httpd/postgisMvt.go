package httpd

import (
	"fmt"
	"github.com/dadadamarine/orb/maptile"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"pbf_demo/internal/logger"
	"pbf_demo/internal/model"
	"pbf_demo/internal/params"
	"pbf_demo/internal/utils/response"
	"strconv"
)

// @Summary postgis
// @Description  postgis mvt
// @version 1.0
// @tags mvt
// @Param z path int true "z"
// @Param x path int true "x"
// @Param y path int true "y"
// @Success 200 {object} []byte "{"code":200,"data": "","msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"bind query err/params error"}"
// @Router /postgis/{z}/{x}/{y} [get]
func PostgisMVTGet(c *gin.Context) {
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
	tile := maptile.New(uint32(x), uint32(y), maptile.Zoom(z))
	b := tile.Bound()
	sqlStr := fmt.Sprintf(`select ST_AsMVT(P, '%v', 4096, 'the_geom') as "mvt" from (
select fid, ST_AsMVTGeom(ST_Transform(the_geom, 3857),ST_Transform(ST_MakeEnvelope(%v,%v,%v,%v, 4326), 3857),
4096, 64, TRUE) the_geom FROM adcode_test) AS P;`, "adcode", b.Min[0], b.Min[1], b.Max[0], b.Max[1])
	pbf := make([]byte, 0)
	row := model.DB.Debug().Raw(sqlStr).Row()
	row.Scan(&pbf)
	c.Header("Content-Length", fmt.Sprintf("%d", len(pbf)))
	c.Data(http.StatusOK, "application/vnd.mapbox-vector-tile", pbf)
	return
}
