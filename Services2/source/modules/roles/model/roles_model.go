// Archivo generado automáticamente para el módulo roles (model)
package roles_model

import (
	user_model "main/connection/services/user_service"

	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
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
