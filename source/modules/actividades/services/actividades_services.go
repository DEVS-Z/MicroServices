package actividades_services

import (
	base_service "main/pkg/base/service"
	actividades_model "main/source/modules/actividades/model"
	"main/source/modules/estadoActividad"
	estado_actividad_model "main/source/modules/estadoActividad/model"
	"main/source/modules/tipoActividad"
	tipo_actividad_model "main/source/modules/tipoActividad/model"
)

type ActividadesService struct {
	base_service.Service[actividades_model.ActividadesStruct]
}

func (s *ActividadesService) Read(filters map[string]interface{}) ([]actividades_model.ActividadesSanitizer, error) {
	var actividadesSanitizadas []actividades_model.ActividadesSanitizer
	actividades, err := s.Service.Read(filters)
	if err != nil {
		return nil, err
	}

	for _, actividad := range actividades {
		if actividad.Tipo == nil || actividad.Estado == nil {
			continue // Omitir si Tipo o Estado son nil
		}
		tipo, err := tipoActividad.Service.ReadOne(*actividad.Tipo)
		if err != nil {
			tipo = tipo_actividad_model.TipoActividadStruct{}
		}

		estado, err := estadoActividad.Service.ReadOne(*actividad.Estado)
		if err != nil {
			estado = estado_actividad_model.EstadoActividadStruct{}
		}

		actividadesSanitizadas = append(actividadesSanitizadas, actividades_model.ActividadesSanitizer{
			ActividadesStruct: actividad,
			Tipo:              &tipo.Nombre,
			Estado:            &estado.Nombre,
		})
	}

	return actividadesSanitizadas, nil
}

func (s *ActividadesService) ReadOne(id int) (*actividades_model.ActividadesSanitizer, error) {
	var tipo tipo_actividad_model.TipoActividadStruct
	var estado estado_actividad_model.EstadoActividadStruct
	actividad, err := s.Service.ReadOne(id)
	if err != nil {
		return nil, err
	}

	if actividad.Tipo != nil || actividad.Estado != nil {

		tipo, err = tipoActividad.Service.ReadOne(*actividad.Tipo)
		if err != nil {
			return nil, err
		}
		estado, err = estadoActividad.Service.ReadOne(*actividad.Estado)
		if err != nil {
			return nil, err
		}
	} else {
		tipo = tipo_actividad_model.TipoActividadStruct{}
		estado = estado_actividad_model.EstadoActividadStruct{}
	}

	return &actividades_model.ActividadesSanitizer{
		ActividadesStruct: actividad,
		Tipo:              &tipo.Nombre,
		Estado:            &estado.Nombre,
	}, nil
}

var Service = base_service.NewService[ActividadesService](*actividades_model.Model)
