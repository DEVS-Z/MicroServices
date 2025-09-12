package router

import (
	"main/pkg/base/router"

	"github.com/gin-contrib/cors"
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
		AllowCredentials: true,
	}))

}
