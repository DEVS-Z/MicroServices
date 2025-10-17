// Archivo generado automáticamente para el módulo Suscripciones (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type SuscripcionesStruct struct {
	SuscripcionId *int     `db:"suscripcion_id" sanitizer:"id" visible:"false"`
	Nombre        *string  `db:"nombre"`
	Precio        *float64 `db:"precio"`
	CantPlayers   *int     `db:"cant_players"`
}

var Model = base_models.NewModel[SuscripcionesStruct]("suscripciones", "zfut.suscripciones")
