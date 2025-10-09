package ejercicios

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	ejercicios_model "main/source/modules/ejercicios/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[ejercicios_model.EjerciciosStruct]](*ejercicios_model.Model)

func Init() {
	print("Ejercicios Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/ejercicios")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
