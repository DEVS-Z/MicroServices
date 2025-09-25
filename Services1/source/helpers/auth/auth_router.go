package auth

import (
	"main/source/helpers/router"
)

func AuthRouter() {
	r := router.NewRoute("/auth")

	r.POST("/sign-in", SignIn)
	r.POST("/sign-up", SignUp)
}
