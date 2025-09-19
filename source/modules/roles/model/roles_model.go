// Archivo generado automáticamente para el módulo roles (model)
package roles_model

import (
	base_models "main/pkg/base/models"
	user_model "main/source/modules/users/models"
)

type RolesStruct struct {
	Id     *int   `db:"id" sanitizer:"id" visible:"false"`
	Nombre string `db:"nombre"`
}

type UsersInRoles struct {
	Role  *string
	Users []user_model.UserSanitizer
}

var Model = base_models.NewModel[RolesStruct]("roles", "roles")
