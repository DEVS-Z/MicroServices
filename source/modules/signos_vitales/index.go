package signos_vitales

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	signos_vitales_model "main/source/modules/signos_vitales/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[signos_vitales_model.Signos_vitalesStruct]](*signos_vitales_model.Model)

func Init() {
	print("Signos_vitales Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/signos_vitales")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
