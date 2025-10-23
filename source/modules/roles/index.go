package roles

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	roles_model "main/source/modules/roles/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[roles_model.RolesStruct]](*roles_model.Model)

func Init() {
	print("Roles Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/roles")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
