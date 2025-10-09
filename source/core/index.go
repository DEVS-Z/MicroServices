package core

import (
	"main/source/modules/actividades"
	"main/source/modules/alertas"
	"main/source/modules/asignacion_rutinas"
	"main/source/modules/auth"
	"main/source/modules/club_suscripcion"
	"main/source/modules/clubs"
	"main/source/modules/ejercicios"
	"main/source/modules/equipos"
	"main/source/modules/eventos"
	"main/source/modules/miembros"
	"main/source/modules/pagos"
	"main/source/modules/reportes"
	"main/source/modules/roles"
	"main/source/modules/rutinas"
	"main/source/modules/signos_vitales"
	"main/source/modules/suscripciones"
	"main/source/modules/usuarios"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func Init() {
	// Aquí se inicializan los módulos

	modules.NewModule(roles.Init)

	modules.NewModule(usuarios.Init)

	modules.NewModule(clubs.Init)

	modules.NewModule(suscripciones.Init)

	modules.NewModule(club_suscripcion.Init)

	modules.NewModule(pagos.Init)

	modules.NewModule(miembros.Init)

	modules.NewModule(equipos.Init)

	modules.NewModule(eventos.Init)

	modules.NewModule(rutinas.Init)

	modules.NewModule(ejercicios.Init)

	modules.NewModule(asignacion_rutinas.Init)

	modules.NewModule(actividades.Init)

	modules.NewModule(signos_vitales.Init)

	modules.NewModule(alertas.Init)

	modules.NewModule(reportes.Init)

	auth.InitAuth()
}
