// Archivo generado automáticamente para el módulo Clubs (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type ClubsStruct struct {
	ClubId        *int    `db:"club_id" sanitizer:"id" visible:"false" type: "pk"`
	Nombre        *string `db:"nombre"`
	Pais          *string `db:"pais"`
	FechaRegistro *string `db:"fecha_registro"`
	OwnerId       *int    `db:"owner_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[ClubsStruct]("clubs", "zfut.clubs")
