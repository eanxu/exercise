package main

import (
	"fmt"
	"github.com/airmap/gdal"
	"github.com/airmap/gdal/ogr"
	"cutTile/getBounds/bounds"
	"cutTile/getBounds/realTime/utils"
	"log"
)

// 栅格实时获取切片边框及中心点

func getTifInformation(path string) bounds.CRSBounds {
	//     val indataset = gdal.Open(path)
	//    val bandcount = indataset.GetRasterCount()
	//    val min:Array[Double] =Array.fill(bandcount)(0)
	//    val max:Array[Double] =Array.fill(bandcount)(255)
	//    val nodatas:Array[java.lang.Double] =Array.fill(bandcount)(null)
	//    val datatype = indataset.GetRasterBand(1).getDataType
	dataset, err := gdal.Open(path, gdal.ReadOnly)
	if err != nil {
		log.Fatal(err)
	}
	bandcount := dataset.RasterCount()
	min := make([]float64, 0)
	max := make([]float64, 0)
	nodatas := make([]float64, 0)
	datatype := dataset.RasterBand(1).RasterDataType()
	for i := 0 ; i < bandcount; i++ {
		min = append(min, 0)
		max = append(max, 255)
		nodatas = append(nodatas, 0)
	}

	if datatype > 1 {
		for iband := 0; iband < bandcount; iband++ {
			band := dataset.RasterBand(iband+1)
			// val bandmin: Array[Double]= new Array[Double](1);
			//        val bandmax: Array[Double]= new Array[Double](1)
			//        val mean: Array[Double]= new Array[Double](1)
			//        val stddev: Array[Double]= new Array[Double](1)
			//        band.GetStatistics(true,true,bandmin,bandmax,mean,stddev)
			minBand, maxBand, _, _ := band.GetStatistics(1, 1)
			//         maxminBand(0)=bandmin(0) -> minBand
			//        maxminBand(1)=bandmax(0) -> maxBand
			//        val srcdis = maxminBand(1) - maxminBand(0)
			//        var srchistbuf = srcdis.doubleValue()/512
			srchistbuf := (maxBand - minBand) / 512
			if srchistbuf > 0.5 {
				srchistbuf = 0.5
			}
			//val pHistogramB = new Array[Int](256)
			//        val bandminval = maxminBand(0)-srchistbuf
			//        val res = band.GetHistogram(bandminval,maxminBand(1)+srchistbuf,pHistogramB,true,true)
			//        val r = CalculateCumulativeCount(pHistogramB,0.02,0.98)
			bandminval := minBand - srchistbuf
			bandmaxval := maxBand + srchistbuf
			p := gdal.ProgressFunc(gdal.DummyProgress)
			var data interface{}
			histogram, err := band.Histogram(bandminval, bandmaxval, 256, 1, 1, p, data)
			if err != nil {
				log.Fatal(err)
			}
			minIndex, maxIndex := calculateCumulativeCount(histogram, 0.02, 0.98)
			//         val dis = maxminBand(1)-bandminval
			//        val dissingle = dis.doubleValue()/256
			//        val minv = maxminBand(0)+dissingle*r._1
			//        val maxv = maxminBand(0)+dissingle*r._2
			//         min(iband)=if(minv>0){
			//          minv
			//        }else{
			//          1
			//        }
			//        max(iband)=maxv
			//         maxminBand(0)=bandmin(0) -> minBand
			//        maxminBand(1)=bandmax(0) -> maxBand
			dissingle := (maxBand - bandminval) / 256
			minv := minBand + dissingle * float64(minIndex)
			maxv := minBand + dissingle * float64(maxIndex)

			if minv > 0 { min[iband] = minv } else { min[iband] = 1 }
			max[iband] = maxv
		}
	} else {
		for iband := 0; iband < bandcount; iband++ {
			band := dataset.RasterBand(iband + 1)
			p := gdal.ProgressFunc(gdal.DummyProgress)
			var data  interface{}
			bandmin, bandmax, _, _ := band.ComputeStatistics(1, p, data)
			min[iband] = bandmin
			max[iband] = bandmax
		}
	}

	for iband := 0; iband < bandcount; iband++ {
		band := dataset.RasterBand(iband + 1)
		nodatavalue, ok := band.NoDataValue()
		if !ok {
			log.Fatal("NoDataValue 不知道有没有问题, 但先搞个报错放着先")
		}
		if nodatavalue != 0 {
			nodatas[iband] = nodatavalue
		}
	}

	inSRS := ogr.CreateSpatialReference("")
	utils.SetupInputSRS(dataset, inSRS)
	geoTrans := dataset.GeoTransform()
	inxsize := dataset.RasterXSize()
	inysize := dataset.RasterYSize()
	orgbounds := bounds.CRSBounds{}
	orgbounds.SetCRSBounds(geoTrans[0], geoTrans[3], geoTrans[0] + (float64(inxsize) * geoTrans[1]), geoTrans[3] + float64(inysize) * geoTrans[5])

	//outputSRS := utils.SetupOutputSRS(inSRS, false)
	//warpedInputDataSet := utils.ReprojectDataset(dataset, inSRS, outputSRS)
	//outGT := warpedInputDataSet.GeoTransform()
	//outxsize := warpedInputDataSet.RasterXSize()
	//outysize := warpedInputDataSet.RasterYSize()
	//prjbounds := bounds.CRSBounds{}
	//prjbounds.SetCRSBounds(outGT[0], outGT[3] + float64(outysize) * outGT[5], outGT[0] + (float64(outxsize) * outGT[1]), outGT[3])
	//warpedInputDataSet.Close()

	dataset.Close()

	fmt.Println("min, ", min)  // 81,279,39,6
	fmt.Println("max, ", max)  // 999，1023，999，1023
	fmt.Println("nodatas, ", nodatas)


	return orgbounds

}

func calculateCumulativeCount(histogram []int, minPercent, maxPercent float64) (int, int) {
	var (
		sum, cursum, minIndex, maxIndex int
	)
	for _, v := range(histogram) {
		sum += v
	}
	length := len(histogram)
	minCellCount := float64(sum) * minPercent
	maxCellCount := float64(sum) * (1 - maxPercent)  //TODO
	maxIndex = length-1
	for i := 0; i < length; i++ {
		cursum += histogram[i]
		if float64(cursum) > minCellCount {
			if i > 0 {
				minIndex = i - 1
			} else {
				minIndex = i
			}
			break
		}
	}
	cursum = 0
	for i := length - 1; i >= 0; i-- {
		cursum += histogram[i]
		if float64(cursum) > maxCellCount {
			maxIndex = i
			break
		}
	}
	return minIndex, maxIndex
}

// setCrsBounds(dstBounds, Math.min(curBounds.getMinX, dstBounds.getMinX), Math.min(curBounds.getMinY, dstBounds.getMinY),
//      Math.max(curBounds.getMaxX, dstBounds.getMaxX), Math.max(curBounds.getMaxY, dstBounds.getMaxY))
//  }

func main() {
	path := `/home/marshmallow/Desktop/vm/vm/temp/cutTile/GF1_PMS2_E116.6_N29.4_20200731_L2E0004961884-MSS2.tiff`
	//paths := []string{
	//	`/home/marshmallow/Desktop/vm/vm/temp/cutTile/GF1C_PMS_E119.4_N35.8_20190831_L1A1021465858-MUX4.tif`,
	//	`/home/marshmallow/Desktop/vm/vm/temp/cutTile/GF1B_PMS_E83.5_N43.0_20201108_L3A1227888051-MUX.tiff`,
	//}
	orgbounds := getTifInformation(path)
	fmt.Println("orgbounds: ", orgbounds)

}
