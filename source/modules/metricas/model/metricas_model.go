// Archivo generado automáticamente para el módulo Signos_vitales (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type MetricasStruct struct {
	MetricaId     *int    `db:"metrica_id" sanitizer:"id" visible:"false" type:"pk"`
	ActividadId   *int    `db:"actividad_id" sanitizer:"id"`
	FechaRegistro *string `db:"fecha_registro" sanitizer:"datetime" visible:"false"`
	Metrica       *string `db:"metrica" sanitizer:"text"`
	Valor         *string `db:"valor" sanitizer:"text"`
	UnidadMedida  *string `db:"unidad_medida" sanitizer:"text"`
}

var Model = base_models.NewModel[MetricasStruct]("metricas", "zfut.metricas")
