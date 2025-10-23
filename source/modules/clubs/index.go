package clubs

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	clubs_model "main/source/modules/clubs/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[clubs_model.ClubsStruct]](*clubs_model.Model)

func Init() {
	print("Clubs Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/clubs")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
