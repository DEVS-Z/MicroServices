package core

import (
	"main/source/modules/actividades"
	"main/source/modules/auth"
	"main/source/modules/ejercicios"
	"main/source/modules/jugador"
	"main/source/modules/jugador_rutinas"
	"main/source/modules/roles"
	"main/source/modules/rutinas"
	"main/source/modules/usuarios"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func Init() {
	// Aquí se inicializan los módulos de la aplicación

	modules.NewModule(roles.Init)

	modules.NewModule(usuarios.Init)

	modules.NewModule(rutinas.Init)

	modules.NewModule(ejercicios.Init)

	modules.NewModule(actividades.Init)

	modules.NewModule(jugador.Init)

	modules.NewModule(jugador_rutinas.Init)

	auth.InitAuth()
}
