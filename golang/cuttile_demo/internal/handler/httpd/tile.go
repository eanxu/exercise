package httpd

import (
	"cuttile_demo/internal/logger"
	"cuttile_demo/internal/params"
	"cuttile_demo/internal/utils/response"
	"cuttile_demo/internal/utils/tiles"
	"cuttile_demo/internal/utils/tiles/bounds"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// @Summary mvt
// @Description  mvt
// @version 1.0
// @tags mvt
// @Param z path int true "z"
// @Param x path int true "x"
// @Param y path int true "y"
// @Success 200 {object} []byte "{"code":200,"data": "","msg":"success"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"bind query err/params error"}"
// @Router /tile/{z}/{x}/{y} [get]
func TileGet(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	p := params.TileGet{}
	if err := c.BindUri(&p); err != nil {
		logger.Logger.Error("bind uri err",
			zap.Error(err))
		utilGin.Response(400, "bind uri err", nil)
		return
	}
	path := `/home/ean/Desktop/vm/vm/temp/cutTile/GF1_PMS2_E116.6_N29.4_20200731_L2E0004961884-MSS2.tiff`
	min := []float64{353.6796875, 353.6796875, 353.6796875, 353.6796875}
	max := []float64{791.40234375, 791.40234375, 791.40234375, 791.40234375}
	nodatas := []float64{0, 0, 0, 0}
	// {116.71875 29.22889003019423 116.806640625 29.15216128331892}
	tiffBound := bounds.Bounds{LeftLon:  116.4043777902222, LeftLat: 29.543272604576725, RightLon: 116.86035049875385, RightLat: 29.150346832417437}
	pngBytes := make([]byte, 0)
	pngBytes, ok := tiles.GetTiles(p.X, p.Y, p.Z, tiffBound, path, min, max, nodatas)
	if ok {
		c.Header("Content-Length", fmt.Sprintf("%d", len(pngBytes)))
		c.Data(http.StatusOK, "image/png", pngBytes)
	} else {
		c.Data(http.StatusOK, "", []byte{})
	}
	return
}
