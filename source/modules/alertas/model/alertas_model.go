// Archivo generado automáticamente para el módulo Alertas (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type AlertasStruct struct {
	AlertasId *int `db:"alertas_id" sanitizer:"id" visible:"false"`
	ActividadId *int `db:"actividad_id" sanitizer:"id" visible:"false"`
	FechaRegistro *string `db:"fecha_registro"`
	Tipo *string `db:"tipo"`
	Descripcion *string `db:"descripcion"`
	Gravedad *string `db:"gravedad"`
	AtendidoSiNo *bool `db:"atendido_si_no"`
	MiembroId *int `db:"miembro_id" sanitizer:"id" visible:"false"`
}

var Model = base_models.NewModel[AlertasStruct]("alertas", "alertas")
