package util

import (
	actividades_model "main/source/modules/actividades/model"
	actividades_services "main/source/modules/actividades/services"
)

func ActividadesByEstado(filtro map[string]interface{}) ([]actividades_model.ActividadesPorEstado, error) {
	// Obtener actividades con el filtro aplicado
	dataSanitizer, err := actividades_services.Service.Read(filtro)
	if err != nil {
		return nil, err
	}

	// Usamos un mapa para agrupar por estado
	agrupadas := make(map[string]*actividades_model.ActividadesPorEstado)
	var sinEstado *actividades_model.ActividadesPorEstado // Para actividades con estado nil

	for _, actividad := range dataSanitizer {
		if actividad.Estado == nil {
			if sinEstado == nil {
				sinEstado = &actividades_model.ActividadesPorEstado{
					Estado:      nil,
					Actividades: []actividades_model.ActividadesSanitizer{},
				}
			}
			sinEstado.Actividades = append(sinEstado.Actividades, actividad)
			continue
		}

		estadoID := *actividad.Estado
		if _, exists := agrupadas[estadoID]; !exists {
			agrupadas[estadoID] = &actividades_model.ActividadesPorEstado{
				Estado:      &estadoID,
				Actividades: []actividades_model.ActividadesSanitizer{},
			}
		}
		agrupadas[estadoID].Actividades = append(agrupadas[estadoID].Actividades, actividad)
	}

	// Convertimos el mapa a slice
	var resultado []actividades_model.ActividadesPorEstado
	if sinEstado != nil {
		resultado = append(resultado, *sinEstado)
	}
	for _, item := range agrupadas {
		resultado = append(resultado, *item)
	}

	return resultado, nil
}
