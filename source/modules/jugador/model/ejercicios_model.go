// Archivo generado automáticamente para el módulo Ejercicios (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type JugadorStruct struct {
	JugadorId     *int    `db:"jugador_id" sanitizer:"id" visible:"false" type: "pk"`
	UsuarioId     *int    `db:"user_id" sanitizer:"id" visible:"false"`
	Posicion      *int    `db:"posicion" sanitizer:"int"`
	Altura        *int    `db:"altura" sanitizer:"int"`
	Peso          *int    `db:"peso" sanitizer:"int"`
	FechaRegistro *string `db:"fecha_registro" sanitizer:"datetime"`
	WearableId    *int    `db:"wearable_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[JugadorStruct]("jugadores", "zfut.jugadores")
