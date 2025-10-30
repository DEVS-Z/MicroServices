// Archivo generado automáticamente para el módulo Club_suscripcion (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type Club_suscripcionStruct struct {
	ClubSubsId    *int    `db:"club_subs_id" sanitizer:"id" visible:"false" type: "pk"`
	ClubId        *int    `db:"club_id" sanitizer:"id" visible:"false"`
	SuscripcionId *int    `db:"suscripcion_id" sanitizer:"id" visible:"false"`
	OwnerId       *int    `db:"owner_id" sanitizer:"id" visible:"false"`
	FechaInicio   *string `db:"fecha_inicio"`
	FechaFin      *string `db:"fecha_fin"`
	Status        *string `db:"status"`
	JugadoresAct  *int    `db:"jugadores_act"`
}

var Model = base_models.NewModel[Club_suscripcionStruct]("club_suscripcion", "zfut.club_suscripcion")
