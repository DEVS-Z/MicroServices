package util

import (
	"fmt"
	roles_model "main/source/modules/roles/model"
	roles_services "main/source/modules/roles/services"
	user_model "main/source/modules/users/models"
	user_service "main/source/modules/users/services"
)

func UsersByRole() ([]roles_model.UsersInRoles, error) {
	var err error
	var rolesArr []roles_model.RolesStruct
	var usersArr []user_model.UserSanitizer

	if rolesArr, err = roles_services.Service.Read(nil); err != nil || len(rolesArr) == 0 {
		return nil, fmt.Errorf("error al leer roles: %w", err)
	}

	if usersArr, err = user_service.Service.Read(nil); err != nil || len(usersArr) == 0 {
		return nil, fmt.Errorf("error al leer usuarios: %w", err)
	}

	dataSanitizer := make([]roles_model.UsersInRoles, 0, len(rolesArr))
	for _, role := range rolesArr {
		var users []user_model.UserSanitizer
		for _, user := range usersArr {
			if user.ID != 0 && *user.Rol == role.Nombre {
				users = append(users, user)
			}
		}
		roleName := role.Nombre // para obtener dirección de una variable concreta
		dataSanitizer = append(dataSanitizer, roles_model.UsersInRoles{
			Role:  &roleName,
			Users: users,
		})
	}

	return dataSanitizer, nil
}
