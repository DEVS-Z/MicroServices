package roles

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	roles_services "main/source/modules/roles/services"
	"main/source/modules/roles/util"
)

func Init() {
	print("Roles Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/roles")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", roles_services.Service.Read)
	r.GET("/users", util.UsersByRole)
	r.POST("/", roles_services.Service.Insert)
	r.GET("/:id", roles_services.Service.ReadOne)
	r.PUT("/:id", roles_services.Service.Update)
	r.DELETE("/:id", roles_services.Service.Delete)
}
