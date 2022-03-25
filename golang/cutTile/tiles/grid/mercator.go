package grid

import (
	"cutTile/tiles/bounds"
	"math"
)

const (
	semiMajorAxis = 6378137
	zoomFactor = 2
	rootTiles = 1
	mercatorSrid = 3857
)

type Mercator struct {
	SemiMajorAxis int64
	EarthCircumference float64
	OriginShift float64
	Grid
}

func (m *Mercator) InitMercator(tileSize int64, yOriginNorth bool) (err error) {
	m.SemiMajorAxis = semiMajorAxis
	m.EarthCircumference = 2 * math.Pi * semiMajorAxis
	m.OriginShift = m.EarthCircumference / 2.0
	b := bounds.Bounds{}
	b.SetBounds(-m.OriginShift, -m.OriginShift, m.OriginShift, m.OriginShift)
	err = m.InitGrid(tileSize, zoomFactor, rootTiles, b, mercatorSrid, yOriginNorth)
	if err != nil {
		return
	}
	return
}

