package tiles

import (
	"cuttile_demo/internal/utils/tiles/bounds"
	"cuttile_demo/internal/utils/tiles/geoQuery"
	"cuttile_demo/internal/utils/tiles/gridGenerate"
	"cuttile_demo/internal/utils/tiles/xyz2lonlat"
	"fmt"
	"github.com/airmap/gdal"
	"log"
)

func GetTiles(x, y, z int64, tileBounds bounds.Bounds, path string, bandminv, bandmaxv, nodatas []float64) ([]byte, bool) {
	g := xyz2lonlat.GridFactory{}
	p := g.Generate("p")
	b := p.XYZ2lonlat(x, y, z)

	if !(b.Intersect(&tileBounds)) {
		//log.Println("没有相交")
		return []byte{}, false
	}

	indata, err := gdal.Open(path, gdal.ReadOnly)
	datatype := indata.RasterBand(1).RasterDataType()
	if err != nil {
		log.Println(err)
		return []byte{}, false
	}
	// longtitude  经度 右 > 左
	// latitude 纬度 上 > 下
	r, w := geoQuery.GeoQuery(indata, b.LeftLon, b.LeftLat, b.RightLon, b.RightLat, 256)
	dstile := gridGenerate.DataSetTile{}
	pngBytes := make([]byte, 0)
	if(r.XSize <= 0 || r.YSize <= 0) {
		indata.Close()
		log.Println(fmt.Sprintf("r.XSize = %v, r.YSize = %v, 不知道是不是无影像错误????", r.XSize, r.YSize))
		return []byte{}, false
	} else {
		// dstile = GridGenerator.CreateTile(warped_input_dataset, datatype, bands, 256, rx, ry, rxsize, rysize, wx, wy, wxsize, wysize,
		//            256, "bilinear", bandminv, bandmaxv, nodatas)

		pngBytes = dstile.CreateTile(indata, datatype, []int{0, 1, 2, 3}, 256, r, w, 256,
			"bilinear", bandminv, bandmaxv, nodatas)
	}
	indata.Close()
	return pngBytes, true
}
