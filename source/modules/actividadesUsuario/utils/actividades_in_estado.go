package util

import (
	actividades_usuario_model "main/source/modules/actividadesUsuario/model"
	actividades_usuarios_service "main/source/modules/actividadesUsuario/services"
)

func ActividadesByEstado(filtro map[string]interface{}) ([]actividades_usuario_model.ActividadesPorEstado, error) {
	data, err := actividades_usuarios_service.Service.Read(filtro)
	if err != nil {
		return nil, err
	}

	agrupadas := make(map[string]*actividades_usuario_model.ActividadesPorEstado)
	var sinEstado *actividades_usuario_model.ActividadesPorEstado

	for _, item := range data {
		if item.Actividad.Estado == nil {
			if sinEstado == nil {
				sinEstado = &actividades_usuario_model.ActividadesPorEstado{
					Estado:      nil,
					Actividades: []actividades_usuario_model.ActividadesUsuarioStructWithRelacion{},
				}
			}
			sinEstado.Actividades = append(sinEstado.Actividades, item)
			continue
		}

		estadoID := *item.Actividad.Estado

		group := agrupadas[estadoID]
		if group == nil {
			group = &actividades_usuario_model.ActividadesPorEstado{
				Estado:      &estadoID,
				Actividades: make([]actividades_usuario_model.ActividadesUsuarioStructWithRelacion, 0),
			}
			agrupadas[estadoID] = group
		}
		group.Actividades = append(group.Actividades, item)
	}

	resultado := make([]actividades_usuario_model.ActividadesPorEstado, 0, len(agrupadas)+1)

	if sinEstado != nil {
		resultado = append(resultado, *sinEstado)
	}

	for _, group := range agrupadas {
		resultado = append(resultado, *group)
	}

	return resultado, nil
}
