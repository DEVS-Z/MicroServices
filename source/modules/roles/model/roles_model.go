// Archivo generado automáticamente para el módulo Roles (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type RolesStruct struct {
	RolId  *int    `db:"rol_id" sanitizer:"id" visible:"false" type: "pk"`
	Nombre *string `db:"nombre"`
}

var Model = base_models.NewModel[RolesStruct]("roles", "zfut.roles")
