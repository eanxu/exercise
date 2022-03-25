package coordinate

// 经纬度的点 - double

type CRSPoint struct {
	X float64
	Y float64
}

func (p *CRSPoint) setX(x float64) { p.X = x }
func (p *CRSPoint) setY(y float64) { p.Y = y }

func (p *CRSPoint) InitCRSPoint(x, y float64) {
	p.setX(x)
	p.setY(y)
}