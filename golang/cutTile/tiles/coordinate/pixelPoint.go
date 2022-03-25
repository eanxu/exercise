package coordinate

// 坐标行列号的点 - 长整型

type PixelPoint struct {
	X int64
	Y int64
}

func (p *PixelPoint) setX(x int64) { p.X = x }
func (p *PixelPoint) setY(y int64) { p.Y = y }

func (p *PixelPoint) InitPixelPoint(x, y int64) {
	p.setX(x)
	p.setY(y)
}
