package xyz2lonlat

import "math"

func XYZ2lonlat(x, y, z int64) [2]float64 {
	n := math.Pow(2, float64(z))
	lonDeg := (float64(x)/n)*360.0 - 180.0
	latRad := math.Atan(math.Sinh(math.Pi * (1 - float64(2*y)/n)))
	latDeg := (180 * latRad) / math.Pi
	return [2]float64{lonDeg, latDeg}
}
