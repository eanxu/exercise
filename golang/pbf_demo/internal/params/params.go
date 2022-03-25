package params

type MVTGet struct {
	Z string `uri:"z" binding:"required"`
	X string `uri:"x" binding:"required"`
	Y string `uri:"y" binding:"required"`
}
