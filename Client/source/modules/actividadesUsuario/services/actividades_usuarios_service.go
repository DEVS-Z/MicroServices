package actividades_usuarios_service

import (
	actividades_services "main/source/modules/actividades/services"
	actividades_usuario_model "main/source/modules/actividadesUsuario/model"
	"main/source/services/users"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

type ActividadesUsuarioService struct {
	base_service.Service[actividades_usuario_model.ActividadesUsuarioStruct]
}

func (s *ActividadesUsuarioService) Read(filters map[string]interface{}) ([]actividades_usuario_model.ActividadesUsuarioStructWithRelacion, error) {

	var result []actividades_usuario_model.ActividadesUsuarioStructWithRelacion
	data, err := s.Service.Read(filters)
	if err != nil {
		return nil, err
	}

	for _, item := range data {
		var newItem actividades_usuario_model.ActividadesUsuarioStructWithRelacion
		actividad, err := actividades_services.Service.ReadOne(*item.Actividad)
		if err != nil {
			return nil, err
		}

		nombreUsuario, err := users.GetUserName(*item.Usuario)
		if err != nil {
			return nil, err
		}

		newItem.Id = item.Id
		newItem.Usuario = &nombreUsuario
		newItem.Actividad = actividad
		result = append(result, newItem)
	}

	return result, nil
}

var Service = base_service.NewService[ActividadesUsuarioService](*actividades_usuario_model.Model)
