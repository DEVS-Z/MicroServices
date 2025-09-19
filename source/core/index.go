package modules

import (
	modules "main/pkg/base/core"
	"main/source/helpers/auth"
	"main/source/modules/actividades"
	adminDB "main/source/modules/adminDB"
	"main/source/modules/modulos"
	modulesRol "main/source/modules/modulosRol"
	"main/source/modules/roles"
	"main/source/modules/users"
)

func init() {
	modules.NewModule(actividades.Init)
	modules.NewModule(modulesRol.Init)
	modules.NewModule(modulos.Init)
	modules.NewModule(roles.Init)
	modules.NewModule(users.Init)
	modules.NewModule(adminDB.Init)
	auth.AuthRouter()

}
