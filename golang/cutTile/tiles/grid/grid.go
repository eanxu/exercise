package grid

import (
	"github.com/airmap/gdal/ogr"
	"cutTile/tiles/bounds"
	"cutTile/tiles/coordinate"
	"math"
)

type Grid struct {
	TileSize          int64
	Bounds            bounds.Bounds
	Srs               ogr.SpatialReference
	InitialResolution float64
	XOriginShift      float64
	YOriginShift      float64
	ZoomFactor        float64
	YOriginNorth      bool  // 起始y在上还是下，google瓦片从左上开始计算  tms从左下计算
}

func (g *Grid)InitGrid(tileSize int64, zoomFactor, rootTiles float64,  b bounds.Bounds, srid int, yOriginNorth bool) (err error) {
	g.TileSize = tileSize
	g.Bounds = b
	g.InitialResolution = (b.GetWidth() / rootTiles) / float64(tileSize)
	g.XOriginShift = b.GetWidth() / 2
	g.YOriginShift = b.GetHeight() / 2
	g.ZoomFactor = zoomFactor

	g.YOriginNorth = yOriginNorth

	// 设置epsg
	g.Srs = ogr.CreateSpatialReference("")
	err = g.Srs.FromEPSG(srid)
	if err != nil {
		return
	}
	g.Srs.SetAxisMappingStrategy(ogr.OAMS_TRADITIONAL_GIS_ORDER)
	return
}

func (g *Grid) resolution(zoom int64) float64 {
	return g.InitialResolution / math.Pow(g.ZoomFactor, float64(zoom))
}

// public CRSPoint pixelsToCrs(final PixelPoint pixel, int zoom) {
//        double res = resolution(zoom);
//        if(!yOriginNorth){
//            return new CRSPoint((pixel.getX() * res) - mXOriginShift, mYOriginShift - (pixel.getY() * res) );
//        }else{
//            return new CRSPoint((pixel.getX() * res) - mXOriginShift, (pixel.getY() * res) - mYOriginShift);
//        }
//
//    }
func (g *Grid) pixelsToCrs(pixel coordinate.PixelPoint, zoom int64) (crs coordinate.CRSPoint) {
	res := g.resolution(zoom)
	if !g.YOriginNorth {
		crs.InitCRSPoint((float64(pixel.X) * res) - g.XOriginShift, g.YOriginShift - (float64(pixel.Y) * res))
		return
	} else {
		crs.InitCRSPoint((float64(pixel.X) * res) - g.XOriginShift, (float64(pixel.Y) * res) - g.YOriginShift)
		return
	}
}


func (g *Grid) TileCrsBounds(coord coordinate.TileCoordinate) (b bounds.CRSBounds) {
	var pxLowerLeft, pxUpperRight coordinate.PixelPoint
	//final PixelPoint pxLowerLeft = new PixelPoint(coord.getX() * mTileSize, coord.getY() * mTileSize);
	pxLowerLeft.InitPixelPoint(coord.X * g.TileSize, coord.Y * g.TileSize)
	//final PixelPoint pxUpperRight = new PixelPoint((coord.getX() + 1) * mTileSize, (coord.getY() + 1) * mTileSize);
	pxUpperRight.InitPixelPoint((coord.X + 1) * g.TileSize, (coord.Y + 1) * g.TileSize)

	var upperLeft coordinate.CRSPoint = g.pixelsToCrs(pxLowerLeft, coord.Z)
	var lowerRight coordinate.CRSPoint = g.pixelsToCrs(pxUpperRight, coord.Z)
	b.SetCRSBounds(upperLeft.X, lowerRight.Y, lowerRight.X, upperLeft.Y)
	return
}

func (g *Grid) SetyOriginNorth(yOriginNorth bool) {
	g.YOriginNorth = yOriginNorth
}
