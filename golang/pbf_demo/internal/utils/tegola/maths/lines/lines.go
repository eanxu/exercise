package lines

import "pbf_demo/internal/utils/tegola"

func FromTLineString(lnstr tegola.LineString) (ln [][2]float64) {
	pts := lnstr.Subpoints()
	for i := range pts {
		ln = append(ln, [2]float64{pts[i].X(), pts[i].Y()})
	}
	return ln
}
