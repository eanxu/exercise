package params

type TileGet struct {
	Z  int64  `uri:"z" binding:"required"`
	X  int64  `uri:"x" binding:"required"`
	Y  int64  `uri:"y" binding:"required"`
}

