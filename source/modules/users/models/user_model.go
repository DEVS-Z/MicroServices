package user_model

import (
	base_models "main/pkg/base/models"
)

type UserStruct struct {
	ID              int     `db:"id"`
	PrimerNombre    string  `db:"primernombre"`
	SegundoNombre   *string `db:"segundonombre"`
	PrimerApellido  string  `db:"primerapellido"`
	SegundoApellido *string `db:"segundoapellido"`
	Matricula       string  `db:"matricula"`
	Correo          string  `db:"correo"`
	Contrasena      string  `db:"contrasena" sanitizer:"contrasena" visible:"false"`
	Rol             int     `db:"rol" sanitizer:"rol" visible:"false"`
}

var Model = base_models.NewModel[UserStruct]("user", "usuarios", 4)
