package restmain

import (
	"fmt"
	"main/source/helpers/router"
)

var version = "1.3.4"

func RunRest() {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("client - REST - Version: %s\n", version)
	Execute()
}

func Execute() {
	r := router.Router()
	r.Execute(":8080")
}
