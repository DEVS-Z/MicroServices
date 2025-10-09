// Archivo generado automáticamente para el módulo Actividades (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type ActividadesStruct struct {
	ActividadId *int `db:"actividad_id" sanitizer:"id" visible:"false"`
	JugadorId *int `db:"jugador_id" sanitizer:"id" visible:"false"`
	MiembroId *int `db:"miembro_id" sanitizer:"id" visible:"false"`
	Tipo *string `db:"tipo"`
	FechaInicio *string `db:"fecha_inicio"`
	FechaFin *string `db:"fecha_fin"`
	Comentarios *string `db:"comentarios"`
}

var Model = base_models.NewModel[ActividadesStruct]("actividades", "actividades")
