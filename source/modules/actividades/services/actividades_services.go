package actividades_services

import (
	actividades_model "main/source/modules/actividades/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

type ActividadesService struct {
	base_service.Service[actividades_model.ActividadesStruct]
}

var Service = base_service.NewService[ActividadesService](*actividades_model.Model)
