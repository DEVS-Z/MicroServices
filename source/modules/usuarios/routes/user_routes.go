package routes

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	user_model "main/source/modules/usuarios/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

func InitRoutes(Service base_service.Default[user_model.UsuariosStruct]) {
	var r = router.NewRoute("/usuarios")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
