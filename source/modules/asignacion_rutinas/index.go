package asignacion_rutinas

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	asignacion_rutinas_model "main/source/modules/asignacion_rutinas/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[asignacion_rutinas_model.Asignacion_rutinasStruct]](*asignacion_rutinas_model.Model)

func Init() {
	print("Asignacion_rutinas Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/asignacion_rutinas")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
