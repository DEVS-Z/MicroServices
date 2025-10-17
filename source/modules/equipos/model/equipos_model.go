// Archivo generado automáticamente para el módulo Equipos (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type EquiposStruct struct {
	EquipoId      *int    `db:"equipo_id" sanitizer:"id" visible:"false"`
	ClubId        *int    `db:"club_id" sanitizer:"id" visible:"false"`
	Nombre        *string `db:"nombre"`
	Categoria     *string `db:"categoria"`
	FechaRegistro *string `db:"fecha_registro"`
}

var Model = base_models.NewModel[EquiposStruct]("equipos", "zfut.equipos")
