package jugador

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	jugador_model "main/source/modules/jugador/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[jugador_model.JugadorStruct]](*jugador_model.Model)

func Init() {
	print("Jugador Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/jugadores")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
