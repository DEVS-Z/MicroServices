package alertas

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	alertas_model "main/source/modules/alertas/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[alertas_model.AlertasStruct]](*alertas_model.Model)

func Init() {
	print("Alertas Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/alertas")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)

	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
