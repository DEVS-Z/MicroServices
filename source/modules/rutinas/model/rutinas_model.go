// Archivo generado automáticamente para el módulo Rutinas (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type RutinasStruct struct {
	RutinaId        *int    `db:"rutina_id" sanitizer:"id" visible:"false" type: "pk"`
	CreadoPorId     *int    `db:"creado_por_id" sanitizer:"id" visible:"false"`
	EquipoId        *int    `db:"equipo_id" sanitizer:"id" visible:"false"`
	Nombre          *string `db:"nombre"`
	Objetivo        *string `db:"objetivo"`
	Tipo            *string `db:"tipo"`
	NivelDificultad *string `db:"nivel_dificultad"`
	FechaRegistro   *string `db:"fecha_registro"`
}

var Model = base_models.NewModel[RutinasStruct]("rutinas", "zfut.rutinas")
