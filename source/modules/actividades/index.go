package actividades

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	actividades_model "main/source/modules/actividades/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[actividades_model.ActividadesStruct]](*actividades_model.Model)

func Init() {
	print("Actividades Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/actividades")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
