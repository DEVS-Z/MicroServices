package modules

import (
	"main/source/helpers/auth"
	"main/source/modules/actividades"
	adminDB "main/source/modules/adminDB"
	"main/source/modules/modulos"
	modulesRol "main/source/modules/modulosRol"
	"main/source/modules/roles"
	"main/source/modules/users"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func Init() {
	print("Cargando módulos...")
	modules.NewModule(actividades.Init)
	modules.NewModule(modulesRol.Init)
	modules.NewModule(modulos.Init)
	modules.NewModule(roles.Init)
	modules.NewModule(users.Init)
	modules.NewModule(adminDB.Init)
	auth.AuthRouter()

}
