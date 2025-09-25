package roles_services

import (
	roles_model "main/source/modules/roles/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

type RoleService struct {
	base_service.Service[roles_model.RolesStruct]
}

var Service = base_service.NewService[base_service.Default[roles_model.RolesStruct]](*roles_model.Model)
