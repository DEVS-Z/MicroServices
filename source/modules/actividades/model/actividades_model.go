// Archivo generado automáticamente para el módulo Actividades (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type ActividadesStruct struct {
	ActividadId *int    `db:"actividad_id" sanitizer:"id" visible:"false" type: "pk"`
	JugadorId   *int    `db:"jugador_id" sanitizer:"id" visible:"false"`
	RutinaId    *int    `db:"rutina_id" sanitizer:"id" visible:"false"`
	Tipo        *string `db:"tipo"`
	FechaInicio *string `db:"fecha_inicio"`
	FechaFin    *string `db:"fecha_fin"`
	Descripcion *string `db:"descripcion"`
}

var Model = base_models.NewModel[ActividadesStruct]("actividades", "zfut.actividades")
