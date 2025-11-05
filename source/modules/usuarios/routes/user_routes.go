package routes

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	user_service "main/source/modules/usuarios/service"
	"main/source/services/users/analizer"
)

func InitRoutes(Service *user_service.ServiceType) {
	var r = router.NewRoute("/usuarios")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
	r.POST("/analizer/", analizer.GetUserName)
}
