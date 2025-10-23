package rutinas

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	rutinas_model "main/source/modules/rutinas/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[rutinas_model.RutinasStruct]](*rutinas_model.Model)

func Init() {
	print("Rutinas Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/rutinas")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
