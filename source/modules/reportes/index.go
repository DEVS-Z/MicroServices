package reportes

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	reportes_model "main/source/modules/reportes/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[reportes_model.ReportesStruct]](*reportes_model.Model)

func Init() {
	print("Reportes Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/reportes")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
