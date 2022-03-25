package main

import (
	"cutTile/tiles/bounds"
	"cutTile/tiles/geoQuery"
	"cutTile/tiles/gridGenerate"
	"cutTile/tiles/xyz2lonlat"
	"fmt"
	"github.com/airmap/gdal"
	"log"
)

func getTiles(x, y, z int64, tileBounds bounds.Bounds, path string, bandminv, bandmaxv, nodatas []float64) {
	//m := grid.Mercator{}
	//err := m.InitMercator(256, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tilept := coordinate.TileCoordinate{}
	//tilept.InitTileCoordinate(x, y, z)
	//
	//var gridtileCrsBounds bounds.CRSBounds = m.TileCrsBounds(tilept)
	//fmt.Println(gridtileCrsBounds)

	g := xyz2lonlat.GridFactory{}
	p := g.Generate("p")
	b := p.XYZ2lonlat(x, y, z)
	fmt.Println(b)

	if !(b.Intersect(&tileBounds)) {
		log.Fatal("没有相交")
	}

	indata, err := gdal.Open(path, gdal.ReadOnly)
	defer indata.Close()
	datatype := indata.RasterBand(1).RasterDataType()
	if err != nil {
		log.Fatal(err)
	}
	// val rbwb = GridGenerator.geo_query(warped_input_dataset, tileCrsBounds.getMinX, tileCrsBounds.getMaxY, tileCrsBounds.getMaxX,
	//        tileCrsBounds.getMinY, 256)
	r, w := geoQuery.GeoQuery(indata, b.LeftLon, b.LeftLat, b.RightLon, b.RightLat, 256)
	dstile := gridGenerate.DataSetTile{}
	if(r.XSize <= 0 || r.YSize <= 0) {
		indata.Close()
		log.Fatal(fmt.Sprintf("r.XSize = %v, r.YSize = %v, 不知道是不是无影像错误????", r.XSize, r.YSize))
	} else {
		// dstile = GridGenerator.CreateTile(warped_input_dataset, datatype, bands, 256, rx, ry, rxsize, rysize, wx, wy, wxsize, wysize,
		//            256, "bilinear", bandminv, bandmaxv, nodatas)

		dstile.CreateTile(indata, datatype, []int{0, 1, 2, 3}, 256, r, w, 256,
		"bilinear", bandminv, bandmaxv, nodatas)
	}

}

func main() {
	var x int64 = 3376
	var y int64 = 1700
	var z int64 = 12
	tileBounds := bounds.Bounds{
		LeftLon:  116.4043777902222,
		LeftLat: 29.543272604576725,
		RightLon: 116.86035049875385,
		RightLat: 29.150346832417437,
	}

	path := `/home/ean/Desktop/vm/vm/temp/cutTile/GF1_PMS2_E116.6_N29.4_20200731_L2E0004961884-MSS2.tiff`
	min := []float64{353.6796875, 353.6796875, 353.6796875, 353.6796875}
	max := []float64{791.40234375, 791.40234375, 791.40234375, 791.40234375}
	nodatas := []float64{0, 0, 0, 0}
	getTiles(x, y, z, tileBounds, path, min, max, nodatas)
}
