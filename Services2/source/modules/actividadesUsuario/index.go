package actividadesUsuario

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	actividades_usuarios_service "main/source/modules/actividadesUsuario/services"
)

func Init() {
	print("ActividadesUsuario Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/actividadesUsuario")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", actividades_usuarios_service.Service.Read)
	r.POST("/", actividades_usuarios_service.Service.Insert)
	r.GET("/:id", actividades_usuarios_service.Service.ReadOne)
	r.PUT("/:id", actividades_usuarios_service.Service.Update)
	r.DELETE("/:id", actividades_usuarios_service.Service.Delete)

}
