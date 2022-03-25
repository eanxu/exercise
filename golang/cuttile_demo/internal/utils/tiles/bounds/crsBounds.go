package bounds

type CRSBounds struct {
	Bound []float64
}

func (b *CRSBounds) SetCRSBounds(minx, miny, maxx, maxy float64) {
	b.Bound = append(b.Bound, minx, miny, maxx, maxy)
}

func (b *CRSBounds) GetMinX() float64   { return b.Bound[0] }
func (b *CRSBounds) GetMinY() float64   { return b.Bound[1] }
func (b *CRSBounds) GetMaxX() float64   { return b.Bound[2] }
func (b *CRSBounds) GetMaxY() float64   { return b.Bound[3] }
func (b *CRSBounds) GetWidth() float64  { return b.GetMaxX() - b.GetMinX() }
func (b *CRSBounds) GetHeight() float64 { return b.GetMaxY() - b.GetMinY() }
