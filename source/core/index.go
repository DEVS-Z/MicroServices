package core

import (
	"fmt"
	"main/source/modules/auth"
	"main/source/modules/modulos"
	"main/source/modules/modulosRol"
	"main/source/modules/roles"
	"main/source/modules/users"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func Init() {
	fmt.Println("Cargando módulos...")
	modules.NewModule(modulosRol.Init)
	modules.NewModule(modulos.Init)
	modules.NewModule(roles.Init)
	modules.NewModule(users.Init)
	auth.InitAuth()

}
