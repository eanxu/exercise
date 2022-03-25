package route

import (
	"github.com/gin-gonic/gin"
	"pbf_demo/internal/handler/httpd"
	"pbf_demo/internal/utils/response"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "pbf_demo/docs"
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
		api.GET("/mvt/:z/:x/:y", httpd.MVTGet)
		api.GET("/pyramid/:z/:x/:y", httpd.PyramidMVTGet)
		api.GET("/orb/:z/:x/:y", httpd.OrbMvtGet)
		api.GET("/postgis/:z/:x/:y", httpd.PostgisMVTGet)
		api.GET("/tegola/:z/:x/:y", httpd.TegolaMVTGet)
	}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
