package gridGenerate

import (
	"cuttile_demo/internal/utils/tiles/geoQuery"
	"fmt"
	"github.com/airmap/gdal"
	"go.uber.org/zap/buffer"
	"image"
	"image/color"
	"image/png"
	"log"
)

type DataSetTile struct {
	DsQuery gdal.Dataset
	IsFullTile bool
}

func (dtile *DataSetTile)CreateTile(ds gdal.Dataset, dsdatatype gdal.DataType, bands []int, tileSize int, r, w geoQuery.GQ,
	querysize int, resampling string, bandmin, bandmax, nodatavalues []float64) []byte {
	//var hasEmptyCell bool
	memDrv, err := gdal.GetDriverByName("MEM")
	if err != nil {
		log.Fatal("GetDriverByName MEM error, ", err)
	}
	alphaband := ds.RasterBand(1).GetMaskBand()
	dsTile := memDrv.Create("", tileSize, tileSize, 4, gdal.Byte, []string{})
	defer dsTile.Close()
	dsBandCount := ds.RasterCount()
	odata := make([]int8, querysize * querysize)
	oalpha := make([]int8, querysize * querysize)
	dtile.DsQuery = memDrv.Create("", querysize, querysize, 4, gdal.Byte, []string{})
	defer dtile.DsQuery.Close()
	newBands := make([]int, 3)
	for i := 0; i < 3; i++ {
		if i >= len(bands) {   // 渲染波段
			newBands[i] = 0
		} else {
			newBands[i] = bands[i]
		}
	}

	for i := 0; i < 3; i++ {
		iReadBand := newBands[i] + 1
		if iReadBand > dsBandCount {
			iReadBand = 1
		}
		if i > dsBandCount - 1 {
			iReadBand = 1
		}
		if dsdatatype > 1 {
			curBandMin := bandmin[iReadBand - 1]
			curBandMax := bandmax[iReadBand - 1]
			dis := curBandMax - curBandMin
			switch dsdatatype {
			case gdal.UInt16, gdal.Int16:
				var buf interface{} = make([]int16, querysize * querysize)
				err := ds.RasterBand(iReadBand).IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, buf, w.XSize, w.YSize, 0, 0)
				if err != nil {
					fmt.Errorf(" gdal.UInt16, gdal.Int16 IO ERROR, ERR:%v", err)
				}
				data := buf.([]int16)
				for icell := 0; icell < len(data); icell++ {
					curvalue := data[icell]
					if len(nodatavalues) >= iReadBand && curvalue == int16(nodatavalues[iReadBand - 1]) {
						//hasEmptyCell = true
						odata[icell] = 0
					} else if curvalue > 0 {
						if float64(curvalue) < curBandMin {
							odata[icell] = 1
						} else if float64(curvalue) > curBandMax {
							odata[icell] = -1
						} else {
							nvalue := int16((float64(curvalue) - curBandMin) * 255 / dis)
							if nvalue < 1 {
								nvalue = 1
							}
							odata[icell] = int8(nvalue)
						}
					} else {
						//hasEmptyCell = true
						odata[icell] = 0
					}
				}
			case gdal.UInt32, gdal.Int32:
				var buf interface{} = make([]int32, querysize * querysize)
				err := ds.RasterBand(iReadBand).IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, buf, w.XSize, w.YSize, 0, 0)
				if err != nil {
					fmt.Errorf(" gdal.UInt32, gdal.Int32 IO ERROR, ERR:%v", err)
				}
				data := buf.([]int32)

				for icell := 0; icell < len(data); icell++ {
					curvalue := data[icell]
					if len(nodatavalues) >= iReadBand && curvalue == int32(nodatavalues[iReadBand - 1]) {
						//hasEmptyCell = true
						odata[icell] = 0
					} else if curvalue > 0 {
						if float64(curvalue) < curBandMin {
							odata[icell] = 1
						} else if float64(curvalue) > curBandMax {
							odata[icell] = -1
						} else {
							nvalue := int32((float64(curvalue) - curBandMin) * 255 / dis)
							if nvalue < 1 {
								nvalue = 1
							}
							odata[icell] = int8(nvalue)
						}
					} else {
						//hasEmptyCell = true
						odata[icell] = 0
					}
				}
			case gdal.Float32:
				var buf interface{} = make([]float32, querysize * querysize)
				err := ds.RasterBand(iReadBand).IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, buf, w.XSize, w.YSize, 0, 0)
				if err != nil {
					fmt.Errorf("gdal.Float32 IO ERROR, ERR:%v", err)
				}
				data := buf.([]float32)

				for icell := 0; icell < len(data); icell++ {
					curvalue := data[icell]

					if curvalue > 0 {
						if float64(curvalue) < curBandMin {
							odata[icell] = 1
						} else if float64(curvalue) > curBandMax {
							odata[icell] = -1
						} else {
							nvalue := int32((float64(curvalue) - curBandMin) * 255 / dis)
							if nvalue < 1 {
								nvalue = 1
							}
							odata[icell] = int8(nvalue)
						}
					} else {
						//hasEmptyCell = true
						odata[icell] = 0
					}
				}
			case gdal.Float64:
				var buf interface{} = make([]float64, querysize * querysize)
				err := ds.RasterBand(iReadBand).IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, buf, w.XSize, w.YSize, 0, 0)
				if err != nil {
					fmt.Errorf("gdal.Float64 IO ERROR, ERR:%v", err)
				}
				data := buf.([]float64)

				for icell := 0; icell < len(data); icell++ {
					curvalue := data[icell]
					if curvalue > 0 {
						if float64(curvalue) < curBandMin {
							odata[icell] = 1
						} else if float64(curvalue) > curBandMax {
							odata[icell] = -1
						} else {
							nvalue := int32((float64(curvalue) - curBandMin) * 255 / dis)
							if nvalue < 1 {
								nvalue = 1
							}
							odata[icell] = int8(nvalue)
						}
					} else {
						//hasEmptyCell = true
						odata[icell] = 0
					}
				}
			default:
				var buf interface{} = make([]int8, querysize * querysize)
				err := ds.RasterBand(iReadBand).IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, buf, w.XSize, w.YSize, 0, 0)
				if err != nil {
					fmt.Errorf("int8 IO ERROR, ERR:%v", err)
				}
				data := buf.([]int8)

				for icell := 0; icell < len(data); icell++ {
					curvalue := data[icell]
					if len(nodatavalues) >= iReadBand && curvalue == int8(nodatavalues[iReadBand - 1]) {
						//hasEmptyCell = true
						odata[icell] = 0
					} else {
						odata[icell] = curvalue
					}
				}
			}
		} else {
			var buf interface{} = make([]int8, querysize * querysize)
			ds.RasterBand(iReadBand).IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, buf, w.XSize, w.YSize, 0, 0)
			data := buf.([]int8)

			for icell := 0; icell < len(data); icell++ {
				curvalue := data[icell]
				if len(nodatavalues) >= iReadBand && curvalue == int8(nodatavalues[iReadBand - 1]) {
					//hasEmptyCell = true
					odata[icell] = 0
				} else {
					odata[icell] = curvalue
				}
			}
		}


		if tileSize == querysize {
			err := dsTile.RasterBand(i+1).IO(gdal.Write, w.X, w.Y, w.XSize, w.YSize, odata, w.XSize, w.YSize, 0, 0)
			if err != nil {
				log.Fatal("tileSize == querysize, IO err, ", err)
			}
		} else {
			err := dtile.DsQuery.RasterBand(i+1).IO(gdal.Write, w.X, w.Y, w.XSize, w.YSize, odata, w.XSize, w.YSize, 0, 0)
			if err != nil {
				log.Fatal("tileSize != querysize, IO err, ", err)
			}
		}
	}

	if tileSize != querysize {
		dtile.DsQuery, dsTile = query2tile(dtile.DsQuery, dsTile, resampling)
	}
	err = alphaband.IO(gdal.Read, r.X, r.Y, r.XSize, r.YSize, oalpha, w.XSize, w.YSize, 0, 0)
	if err != nil {
		log.Fatal("alphaband.IO err, ", err)
	}
	if tileSize == querysize {
		err = dsTile.RasterBand(4).IO(gdal.Write, w.X, w.Y, w.XSize, w.YSize, oalpha, w.XSize, w.YSize, 0, 0)
		if err != nil {
			log.Fatal("tileSize == querysize dtile.DsTile.RasterBand(4).IO err, ", err)
		}
	} else {
		err = dtile.DsQuery.RasterBand(4).IO(gdal.Write, w.X, w.Y, w.XSize, w.YSize, oalpha, w.XSize, w.YSize, 0, 0)
		if err != nil {
			log.Fatal("tileSize != querysize dsquery.RasterBand(4).IO err, ", err)
		}
		dtile.DsQuery, dsTile = query2tile(dtile.DsQuery, dsTile, resampling)
	}

	newb := make([][]uint8, 0)
	for i := 1; i <= 4; i++ {
		var buf interface{} = make([]uint8, 256*256)
		err = dsTile.RasterBand(i).IO(gdal.Read, 0, 0, 256, 256, buf, 256, 256, 0, 0)
		if err != nil {
			log.Fatal("read RasterBand error, err: ", err)
		}
		newb = append(newb, buf.([]uint8))
	}
	p := image.NewRGBA(image.Rect(0, 0, 256, 256))
	n := 0
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			p.Set(x, y, color.RGBA{uint8(newb[0][n]), uint8(newb[1][n]), uint8(newb[2][n]), uint8(newb[3][n])})
			n += 1
		}
	}
	pngBytes := buffer.Buffer{}
	err = png.Encode(&pngBytes, p)
	if err != nil {
		log.Fatal("png encode error, err: ", err)
	}
	return pngBytes.Bytes()
}
