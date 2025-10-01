package modulos

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	modulos_model "main/source/modules/modulos/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[modulos_model.ModulosStruct]](*modulos_model.Model)

func Init() {
	print("Permisos Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/modulos")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
