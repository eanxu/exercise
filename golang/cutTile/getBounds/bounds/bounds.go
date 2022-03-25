package bounds


type Bounds struct {
	LeftLon, LeftLat, RightLon, RightLat float64
}

func (b *Bounds) SetBounds(minx, miny, maxx, maxy float64)  {
	b.LeftLon = minx
	b.LeftLat = miny
	b.RightLon = maxx
	b.RightLat = maxy
}

func (b *Bounds) GetMinX() float64   { return b.LeftLon }
func (b *Bounds) GetMinY() float64   { return b.LeftLat }
func (b *Bounds) GetMaxX() float64   { return b.RightLon }
func (b *Bounds) GetMaxY() float64   { return b.RightLat }
func (b *Bounds) GetWidth() float64  { return b.GetMaxX() - b.GetMinX() }
func (b *Bounds) GetHeight() float64 { return b.GetMaxY() - b.GetMinY() }

