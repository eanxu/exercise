package geoQuery

import (
	"github.com/airmap/gdal"
	"math"
	"sync"
)

type GQ struct {
	X, Y, XSize, YSize int
}

func GeoQuery(ds gdal.Dataset, ulx, uly, lrx, lry float64, querySize int) (r, w GQ) {
	geotran := ds.GeoTransform()
	r.X = int((ulx - geotran[0]) / geotran[1] + 0.001)
	r.Y = int((uly - geotran[3]) / geotran[5] + 0.001)
	r.XSize = int((lrx - ulx) / geotran[1] + 0.5)
	r.YSize = int((lry - uly) / geotran[5] + 0.5)

	if querySize == 0 {
		w.XSize = r.XSize
		w.YSize = r.YSize
	} else {
		w.XSize = querySize
		w.YSize = querySize
	}


	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if r.X < 0 {
			rxshift := int(math.Abs(float64(r.X)))
			w.X = w.XSize * rxshift / r.XSize
			w.XSize = w.XSize - w.X
			r.XSize = r.XSize - r.XSize * rxshift / r.XSize
			r.X = 0
		}
		rasterXSize := ds.RasterXSize()
		if (r.X + r.XSize) > rasterXSize {
			w.XSize = w.XSize * (rasterXSize - r.X) / r.XSize
			r.XSize = rasterXSize - r.X
		}
	}()
	go func() {
		defer wg.Done()
		if r.Y < 0 {
			ryshift := math.Abs(float64(r.Y))
			w.Y = int(float64(w.YSize) * ryshift / float64(r.YSize))
			w.YSize = w.YSize - w.Y
			r.YSize = r.YSize - int(float64(r.YSize) * ryshift / float64(r.YSize))
			r.Y = 0
		}
		rasterYSize := ds.RasterYSize()
		if (r.Y + r.YSize) > rasterYSize {
			w.YSize = w.YSize * (rasterYSize - r.Y) / r.YSize
			r.YSize = rasterYSize - r.Y
		}
	}()
	wg.Wait()
	return
}
