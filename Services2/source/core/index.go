package core

import (
	"fmt"
	"main/source/modules/actividades"
	adminDB "main/source/modules/adminDB"
	"main/source/modules/modulos"
	modulesRol "main/source/modules/modulosRol"
	"main/source/modules/roles"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func Init() {
	fmt.Println("Cargando módulos...")
	modules.NewModule(actividades.Init)
	modules.NewModule(modulesRol.Init)
	modules.NewModule(modulos.Init)
	modules.NewModule(roles.Init)
	modules.NewModule(adminDB.Init)

}
