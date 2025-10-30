// Archivo generado automáticamente para el módulo Ejercicios (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type EjerciciosStruct struct {
	EjercicioId  *int    `db:"ejercicio_id" sanitizer:"id" visible:"false" type: "pk"`
	RutinaId     *int    `db:"rutina_id" sanitizer:"id" visible:"false"`
	Nombre       *string `db:"nombre"`
	Series       *int    `db:"series"`
	Repeticiones *int    `db:"repeticiones"`
	DuracionSegs *int    `db:"duracion_segs"`
	Intensidad   *string `db:"intensidad"`
}

var Model = base_models.NewModel[EjerciciosStruct]("ejercicios", "zfut.ejercicios")
