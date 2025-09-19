package roles_services

import (
	base_service "main/pkg/base/service"
	roles_model "main/source/modules/roles/model"
)

type RoleService struct {
	base_service.Service[roles_model.RolesStruct]
}

var Service = base_service.NewService[base_service.Default[roles_model.RolesStruct]](*roles_model.Model)
