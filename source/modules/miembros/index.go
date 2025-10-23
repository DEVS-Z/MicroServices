package miembros

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	miembros_model "main/source/modules/miembros/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[miembros_model.MiembrosStruct]](*miembros_model.Model)

func Init() {
	print("Miembros Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/miembros")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
