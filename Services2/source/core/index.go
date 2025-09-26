package core

import (
	"fmt"
	"main/source/modules/actividades"
	"main/source/modules/actividadesUsuario"
	adminDB "main/source/modules/adminDB"
	"main/source/modules/modulos"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func Init() {
	fmt.Println("Cargando módulos...")
	modules.NewModule(actividades.Init)
	modules.NewModule(modulos.Init)
	modules.NewModule(adminDB.Init)
	modules.NewModule(actividadesUsuario.Init)

}
