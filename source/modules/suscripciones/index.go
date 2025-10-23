package suscripciones

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	suscripciones_model "main/source/modules/suscripciones/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[suscripciones_model.SuscripcionesStruct]](*suscripciones_model.Model)

func Init() {
	print("Suscripciones Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/suscripciones")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
