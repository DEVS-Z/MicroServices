// Archivo generado automáticamente para el módulo Asignacion_rutinas (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type Asignacion_rutinasStruct struct {
	AsignacionId *int    `db:"asignacion_id" sanitizer:"id" visible:"false" type: "pk"`
	RutinaId     *int    `db:"rutina_id" sanitizer:"id" visible:"false"`
	JugadorId    *int    `db:"jugador_id" sanitizer:"id" visible:"false"`
	EquipoId     *int    `db:"equipo_id" sanitizer:"id" visible:"false"`
	FechaInicio  *string `db:"fecha_inicio"`
	FechaFin     *string `db:"fecha_fin"`
	Frecuencia   *string `db:"frecuencia"`
	Comentarios  *string `db:"comentarios"`
	RolId        *int    `db:"rol_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[Asignacion_rutinasStruct]("asignacion_rutinas", "zfut.asignacion_rutinas")
