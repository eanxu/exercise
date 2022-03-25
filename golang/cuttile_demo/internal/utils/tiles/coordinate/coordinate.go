package coordinate

type coordinate struct {
	X int64
	Y int64
}

func (c *coordinate) setX(x int64) { c.X = x }
func (c *coordinate) setY(y int64) { c.Y = y }

func (c *coordinate) InitCoordinate(x, y int64) {
	c.setX(x)
	c.setY(y)
}