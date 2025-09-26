package auth

import (
	"main/source/helpers/router"
)

func InitAuth() {
	r := router.NewRoute("/auth")

	r.POST("/sign-in", SignIn)
	r.POST("/sign-up", SignUp)
}
