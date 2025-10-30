// Archivo generado automáticamente para el módulo Signos_vitales (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type Signos_vitalesStruct struct {
	SignosId      *int     `db:"signos_id" sanitizer:"id" visible:"false" type: "pk"`
	ActividadId   *int     `db:"actividad_id" sanitizer:"id" visible:"false"`
	FechaRegistro *string  `db:"fecha_registro"`
	Metrica       *string  `db:"metrica"`
	Valor         *float64 `db:"valor"`
	UnidadMedida  *string  `db:"unidad_medida"`
	MiembroId     *int     `db:"miembro_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[Signos_vitalesStruct]("signos_vitales", "zfut.signos_vitales")
