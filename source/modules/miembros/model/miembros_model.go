// Archivo generado automáticamente para el módulo Miembros (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type MiembrosStruct struct {
	MiembroId *int `db:"miembro_id" sanitizer:"id" visible:"false"`
	Posicion *string `db:"posicion"`
	WeareableId *string `db:"weareable_id" sanitizer:"id" visible:"false"`
	FechaRegistro *string `db:"fecha_registro"`
	Altura *float64 `db:"altura"`
	Peso *float64 `db:"peso"`
	UserId *int `db:"user_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[MiembrosStruct]("miembros", "miembros")
