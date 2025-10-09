package equipos

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	equipos_model "main/source/modules/equipos/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[equipos_model.EquiposStruct]](*equipos_model.Model)

func Init() {
	print("Equipos Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/equipos")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
