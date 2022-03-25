package coordinate

type TileCoordinate struct {
	Z int64
	coordinate
}

func (t *TileCoordinate) InitTileCoordinate(x, y, z int64)  {
	t.InitCoordinate(x, y)
	t.Z = z
}
