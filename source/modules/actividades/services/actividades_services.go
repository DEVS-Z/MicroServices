package actividades_services

import (
	base_service "main/pkg/base/service"
	actividades_model "main/source/modules/actividades/model"
)

type ActividadesService struct {
	base_service.Service[actividades_model.ActividadesStruct]
}

var Service = base_service.NewService[ActividadesService](*actividades_model.Model)
