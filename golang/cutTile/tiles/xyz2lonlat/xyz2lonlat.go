package xyz2lonlat

import (
	"cutTile/tiles/bounds"
	"math"
	"sync"
)


type GridFactory struct {}

type GridProduct interface {
	XYZ2lonlat(x, y, z int64) bounds.Bounds
}

func (f GridFactory) Generate(coord string) GridProduct {
	switch coord {
	case "p":
		return ProjProduct{}
	case "g":
		return GEOProduct{}
	default:
		return nil
	}
}

type ProjProduct struct {}

func (p ProjProduct) XYZ2lonlat(x, y, z int64) (b bounds.Bounds) {
	// 投影坐标系
	n := math.Pow(2, float64(z))
	lx := x
	ly := y
	rx := x + 1
	ry := y + 1
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		b.LeftLon = (float64(lx)/n)*360.0 - 180.0
		b.LeftLat = (180 * (math.Atan(math.Sinh(math.Pi * (1 - float64(2*ly)/n))))) / math.Pi
	}()
	go func() {
		defer wg.Done()
		b.RightLon = (float64(rx)/n)*360.0 - 180.0
		b.RightLat = (180 * (math.Atan(math.Sinh(math.Pi * (1 - float64(2*ry)/n))))) / math.Pi
	}()
	wg.Wait()
	return
}

type GEOProduct struct {}

func (g GEOProduct) XYZ2lonlat(x, y, z int64) (b bounds.Bounds) {
	// tms 地理坐标
	n := math.Pow(2, float64(z))
	m := math.Pow(2, float64(z)+1)
	lx := x
	ly := y
	rx := x + 1
	ry := y + 1
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		b.LeftLon = float64(lx)*360.0/m - 180
		b.LeftLat = 90 - float64(ly)*180.0/n
	}()
	go func() {
		defer wg.Done()
		b.RightLon = float64(rx)*360.0/m - 180
		b.RightLat = 90 - float64(ry)*180.0/n
	}()
	wg.Wait()
	return
}