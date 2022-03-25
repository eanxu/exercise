package main

import (
	"bytes"
	"fmt"
	"github.com/airmap/gdal"
	_ "golang.org/x/image/tiff"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func genImage() {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))

	draw.Draw(m, m.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

	f, err := os.Create("./demo.jpeg")
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(f, m, nil)
	if err != nil {
		panic(err)
	}
}

func demo() {
	path := `/home/ean/Desktop/aaa11111.tif`
	ds, err := gdal.Open(path, gdal.ReadOnly)
	defer ds.Close()
	if err != nil {
		log.Fatal("open error, err: ", err)
	}
	var buf interface{} = make([]int16, 256 * 256)
	for i := 1; i <= 4; i++ {
		rd := ds.RasterBand(2)
		err := rd.IO(gdal.Read, 1, 1, 255, 255, buf, 256, 256, 0, 0)
		if err != nil {
			log.Fatal("read error, err: ", err)
		}
		data := buf.([]byte)
		fmt.Println(data)
	}
}

func demo1() {
	path := `/home/ean/Desktop/benchGray.png`
	ds, err := gdal.Open(path, gdal.ReadOnly)
	if err != nil {
		log.Fatal("read error, err: ", err)
	}
	bands := make([][]int32, 0)
	fmt.Println(ds.RasterCount())
	for i := 1; i < 4; i++ {
		var buf interface{} = make([]int32, 256*256)
		err = ds.RasterBand(i).IO(gdal.Read, 0, 0, 256, 256, buf, 256, 256,0, 0)
		if err != nil {
			log.Fatal("read band 1 error, err: ", err)
		}
		data := buf.([]int32)
		bands = append(bands, data)
	}
	p := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			p.Set(x, y, color.RGBA{uint8(bands[0][x*y]), uint8(bands[1][x*y]), uint8(bands[2][x*y]), 0})
		}
	}
	file, err := os.Create("/home/ean/Desktop/vm/vm/temp/cutTile/test/tt1111.png")
	defer file.Close()
	err = png.Encode(file, p)
	if err != nil {
		log.Fatal(err)
	}
}

func demo2() {
	path := `/home/ean/Desktop/benchGray.png`
	f, _ := os.Open(path)
	defer f.Close()
	imageP, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	p := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			r, g, b, a := imageP.At(x,y).RGBA()
			p.Set(x, y, color.RGBA{
				R: uint8(r),
				G: uint8(g),
				B: uint8(b),
				A: uint8(a),
			})
		}
	}
	bb := bytes.Buffer{}
	err = png.Encode(&bb, p)
	bb.Bytes()
}

func demo3() {
	file, err := os.Create("/home/ean/Desktop/vm/vm/temp/cutTile/test/asdf789.png")
	if err != nil {
		log.Fatal("os create error, err: ", err)
	}
	defer file.Close()
	p := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			p.Set(x, y, color.RGBA{
				R: 255,
				G: 50,
				B: 50,
				A: 255,
			})
		}
	}
	err = png.Encode(file, p)
	if err != nil {
		log.Fatal("png encode error, err: ", err)
	}
}

func main() {
	//genImage()
	//demo1()
	//demo2()
	//demo3()
	fmt.Println((116.4043777902222+116.86035049875385)/2)
	fmt.Println((29.543272604576725+29.150346832417437)/2)
}
