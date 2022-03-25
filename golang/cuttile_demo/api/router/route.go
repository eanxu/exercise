package route

import (
	"cuttile_demo/internal/handler/httpd"
	"cuttile_demo/internal/utils/response"
	"github.com/gin-gonic/gin"

	_ "cuttile_demo/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(engine *gin.Engine) {

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	api := engine.Group("/test")
	{

		api.GET("/ping", func(c *gin.Context) {
			utilGin := response.Gin{Ctx: c}
			utilGin.Response(1, "pong", nil)
		})
		api.GET("/tile/:z/:x/:y", httpd.TileGet)
	}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
