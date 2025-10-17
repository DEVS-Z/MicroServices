// Archivo generado automáticamente para el módulo Usuarios (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type UsuariosStruct struct {
	UserId        *int    `db:"user_id" sanitizer:"id" visible:"false"`
	Nombre        *string `db:"nombre"`
	Correo        *string `db:"correo"`
	Password      *string `db:"password"`
	FechaRegistro *string `db:"fecha_registro"`
	Estado        *string `db:"estado"`
	RolId         *int    `db:"rol_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[UsuariosStruct]("usuarios", "zfut.usuarios")
