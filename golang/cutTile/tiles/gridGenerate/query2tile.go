package gridGenerate

import (
	"fmt"
	"github.com/airmap/gdal"
	"log"
)

func query2tile(dsquery, dstile gdal.Dataset, resampling string) (gdal.Dataset, gdal.Dataset) {
	querysize := dsquery.RasterXSize()
	tilesize := dstile.RasterXSize()
	tilebands := dstile.RasterCount()
	if resampling == "average" {
		for i := 1; i <= tilebands; i++ {
			//int res = gdal.RegenerateOverview(dsquery.GetRasterBand(i), dstile.GetRasterBand(i), "average");
			p := gdal.ProgressFunc(gdal.TermProgress)
			var data interface{}
			pahOvrBands := []gdal.RasterBand{dstile.RasterBand(i)}
			res := gdal.RegenerateOverviews(dsquery.RasterBand(i), 1, pahOvrBands,
				"AVERAGE", p, data)
			fmt.Printf("average: res = %v, i = %v \n", res, i)
		}
	} else {
		var gdalResampling gdal.ResampleAlg
		switch resampling {
		case "near":
			gdalResampling = gdal.GRA_NearestNeighbour
		case "bilinear":
			gdalResampling = gdal.GRA_Bilinear
		case "cubic":
			gdalResampling = gdal.GRA_Cubic
		case "cubicspline":
			gdalResampling = gdal.GRA_CubicSpline
		case "lanczos":
			gdalResampling = gdal.GRA_Lanczos
		}
		err := dsquery.SetGeoTransform([6]float64{0, float64(tilesize / querysize), 0, 0, 0, float64(tilesize / querysize)})
		if err != nil {
			log.Fatal("dsquery.SetGeoTransform err, ", err)
		}
		err = dstile.SetGeoTransform([6]float64{0, 1, 0, 0, 0, 1})
		if err != nil {
			log.Fatal("dstile.SetGeoTransform err, ", err)
		}
		//int res = gdal.ReprojectImage(dsquery, dstile, null, null, gdal_resampling);
		p := gdal.ProgressFunc(gdal.DummyProgress)
		var data interface{}
		err = gdal.ReprojectImage(dsquery, dsquery.Projection(), dstile, dsquery.Projection(), gdalResampling, 0, 0, p, data)
		if err != nil {
			log.Fatal("failed to ReprojectImage: %v", err)
		}
	}
	return dsquery, dstile
}
