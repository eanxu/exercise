package bounds

import "fmt"

type Bounds struct {
	LeftLon, LeftLat, RightLon, RightLat float64
}

func (b *Bounds) SetBounds(minx, miny, maxx, maxy float64) {
	b.LeftLon = minx
	b.LeftLat = miny
	b.RightLon = maxx
	b.RightLat = maxy
}

func (b *Bounds) Intersect(otherBound *Bounds) bool {
	fmt.Println("b, ", b)
	fmt.Println("otherBound, ", otherBound)
	if b.LeftLon < otherBound.RightLon && otherBound.LeftLon < b.RightLon &&
		b.RightLat < otherBound.LeftLat && otherBound.RightLat < b.LeftLat {
		return true
	}
	return false
}
