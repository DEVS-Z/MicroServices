package router

import (
	_ "main/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/miqueaz/FrameGo/pkg/base/router"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var rout = router.Router()

func NewRoute(path string) *router.GroupRouter {
	return rout.Group(path)
}

func Router() *router.AppRouter {
	return rout
}

func init() {
	rout.SetTrustedProxies([]string{"*"})
	rout.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
	}))
	rout.Use(gin.Logger())
	rout.Use(gin.Recovery())
	rout.GET("/", func(c *gin.Context) {
		//Redirecciona a swagger
		c.Redirect(302, "/swagger/index.html")
	})
	rout.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

}
