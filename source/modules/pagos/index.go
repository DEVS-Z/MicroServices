package pagos

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	pagos_model "main/source/modules/pagos/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[pagos_model.PagosStruct]](*pagos_model.Model)

func Init() {
	print("Pagos Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/pagos")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
