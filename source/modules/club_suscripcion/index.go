package club_suscripcion

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	club_suscripcion_model "main/source/modules/club_suscripcion/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[club_suscripcion_model.Club_suscripcionStruct]](*club_suscripcion_model.Model)

func Init() {
	print("Club_suscripcion Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/club_suscripcion")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
