package service

import (
	"main/source/modules/usuarios/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

type ServiceType struct {
	base_service.Service[model.UsuariosStruct]
}
