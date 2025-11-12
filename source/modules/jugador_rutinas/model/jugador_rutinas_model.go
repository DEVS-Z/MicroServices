// Archivo generado automáticamente para el módulo Asignacion_rutinas (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type Jugador_rutinasStruct struct {
	AsignacionId    *int    `db:"asignacion_id" sanitizer:"id" visible:"false" type: "pk"`
	RutinaId        *int    `db:"rutina_id" sanitizer:"id" visible:"true" type: "fk"`
	JugadorId       *int    `db:"jugador_id" sanitizer:"id" visible:"true" type: "fk"`
	FechaAsignacion *string `db:"fecha_asignacion" sanitizer:"date" visible:"true" type: "date"`
	Comentarios     *string `db:"comentarios" sanitizer:"text" visible:"true" type: "text"`
}

var Model = base_models.NewModel[Jugador_rutinasStruct]("jugador_rutinas", "zfut.jugador_rutinas")
