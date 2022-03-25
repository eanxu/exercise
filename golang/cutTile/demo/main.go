package main

import "fmt"

//func TestReprojectImage(t *testing.T) {
//	path := "/home/marshmallow/Desktop/vm/vm/temp/cutTile/GF1_PMS2_E116.6_N29.4_20200731_L2E0004961884-MSS2.tiff"
//	ds, err := Open(path, ReadOnly)
//	if err != nil {
//		t.Fatalf("failed to open test file: %v", err)
//	}
//	memDrv, err := GetDriverByName("GTIFF")
//	if err != nil {
//		t.Fatalf("failed to GetDriverByName: %v", err)
//	}
//	dataType := ds.RasterBand(1).RasterDataType()
//	fmt.Println("dataType = ", dataType)
//	fmt.Println("dataType.size = ", dataType.Size())
//	dstPath := `/home/marshmallow/Desktop/vm/vm/temp/cutTile/test/t_go.tiff`
//	dstile := memDrv.Create(dstPath, ds.RasterXSize(), ds.RasterYSize(), 4, dataType, []string{})
//	err = dstile.SetGeoTransform(ds.GeoTransform())
//	if err != nil {
//		t.Fatalf("failed to SetGeoTransform: %v", err)
//	}
//	err = dstile.SetProjection(ds.Projection())
//	if err != nil {
//		t.Fatalf("failed to SetProjection: %v", err)
//	}
//	p := ProgressFunc(DummyProgress)
//	var data interface{}
//	err = ReprojectImage(ds, "", dstile, "", GRA_Bilinear, 0, 0, p, data)
//	if err != nil {
//		t.Fatalf("failed to ReprojectImage: %v", err)
//	}
//	dstileNew, _ := Open(dstPath, ReadOnly)
//	fmt.Println(dstileNew.RasterBand(1).GetMaximum())
//}

func main() {
	//path := `/home/marshmallow/Desktop/vm/vm/temp/cutTile/GF1C_PMS_E119.4_N35.8_20190831_L1A1021465858-MUX4.tif`
	//ds, _ := gdal.Open(path, gdal.ReadOnly)
	//var buf interface{} = make([]int16, 256*256)
	//fmt.Println(ds.RasterCount())
	//ds.RasterBand(1).IO(gdal.Read, 0, 0, 22275, 17901, buf, 256, 256, 0, 0)
	////fmt.Println("buf, ", buf)
	//switch buf.(type) {
	//case []int16:
	//	value := buf.([]int16)
	//	fmt.Println(reflect.TypeOf(value))
	//}
	//a := make([]int, 5)
	//b := make([]int, len(a)+1)
	//fmt.Println(b)
	//b[len(a)] = 111
	//fmt.Println(a)
	//fmt.Println(b)
	a := true
	b := !a
	fmt.Println(b)
}
