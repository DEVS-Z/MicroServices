package actividades_usuarios_service

import (
	base_service "main/pkg/base/service"
	actividades_services "main/source/modules/actividades/services"
	actividades_usuario_model "main/source/modules/actividadesUsuario/model"
)

type ActividadesUsuarioService struct {
	base_service.Service[actividades_usuario_model.ActividadesUsuarioStruct]
}

func (s *ActividadesUsuarioService) Read(filtro map[string]interface{}) ([]actividades_usuario_model.ActividadesUsuarioStructWithRelacion, error) {
	data, err := s.Service.Read(filtro)
	if err != nil {
		return nil, err
	}

	// Convertir a ActividadesUsuarioStructWithRelacion
	var result []actividades_usuario_model.ActividadesUsuarioStructWithRelacion

	for _, item := range data {
		if item.Actividad == nil {
			continue // Omitir si Actividad es nil
		}
		act, err := actividades_services.Service.ReadOne(*item.Actividad)
		if err != nil {
			continue
		}
		result = append(result, actividades_usuario_model.ActividadesUsuarioStructWithRelacion{
			Id:        item.Id,
			Usuario:   item.Usuario,
			Actividad: *act,
		})
	}

	return result, nil
}

var Service = base_service.NewService[ActividadesUsuarioService](*actividades_usuario_model.Model)
