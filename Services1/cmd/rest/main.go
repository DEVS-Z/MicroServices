package restmain

import (
	"fmt"
	config "main/config"
	"main/source/helpers/router"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

var version = "1.3.4"

func RunRest() {
	fmt.Print("\033[H\033[2J")
	config.Execute()
	modules.Execute()
	fmt.Printf("Service1 - REST - Version: %s\n", version)
	Execute()
}

func Execute() {
	r := router.Router()
	r.Execute(":8080")
}
