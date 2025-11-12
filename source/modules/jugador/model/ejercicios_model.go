// Archivo generado automáticamente para el módulo Ejercicios (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type JugadorStruct struct {
	JugadorId    *int    `db:"jugador_id" sanitizer:"id" visible:"false" type: "pk"`
	Nombre       *string `db:"nombre" sanitizer:"text" visible:"true" type: "text"`
	Series       *int    `db:"series" sanitizer:"int" visible:"true" type: "int"`
	Repeticiones *int    `db:"repeticiones" sanitizer:"int" visible:"true" type: "int"`
	DuracionSegs *int    `db:"duracion_segs" sanitizer:"int" visible:"true" type: "int"`
	Intensidad   *string `db:"intensidad" sanitizer:"text" visible:"true" type: "text"`
}

var Model = base_models.NewModel[JugadorStruct]("jugadores", "zfut.jugadores")
