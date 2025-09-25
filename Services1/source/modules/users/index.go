package users

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	my_data_middleware "main/source/helpers/middlewares/myData"
	role_middleware "main/source/helpers/middlewares/rol"
	"main/source/helpers/router"
	user_model "main/source/modules/users/models"
	user_service "main/source/modules/users/services"
)

func Init() {
	print("Users Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/users")
	r.USE(jwt_middleware.JWTMiddleware())
	r.USE(role_middleware.RoleAccessMiddleware(*user_model.Model))
	r.GET("/", user_service.Service.Read)
	r.POST("/", user_service.Service.Insert)
	r.GET("/:id", user_service.Service.ReadOne)
	r.PUT("/:id", user_service.Service.Update)
	r.DELETE("/:id", user_service.Service.Delete)

	var my = router.NewRoute("/users/mis")
	my.USE(jwt_middleware.JWTMiddleware())
	my.USE(my_data_middleware.InjectUserIDAsParamMiddleware())
	my.GET("/", user_service.Service.ReadOne)
	my.PUT("/", user_service.Service.Update)
	my.DELETE("/", user_service.Service.Delete)
}
