package actividades

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	actividades_services "main/source/modules/actividades/services"
	"main/source/modules/actividades/util"
)

func Init() {
	print("Actividades Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/actividades")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", actividades_services.Service.Read)
	r.GET("/estado", util.ActividadesByEstado)
	r.POST("/", actividades_services.Service.Insert)
	r.GET("/:id", actividades_services.Service.ReadOne)
	r.PUT("/:id", actividades_services.Service.Update)
	r.DELETE("/:id", actividades_services.Service.Delete)
}
