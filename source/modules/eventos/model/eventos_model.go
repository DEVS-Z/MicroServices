// Archivo generado automáticamente para el módulo Eventos (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type EventosStruct struct {
	EventoId *int `db:"evento_id" sanitizer:"id" visible:"false"`
	EquipoId *int `db:"equipo_id" sanitizer:"id" visible:"false"`
	CreadoPorId *int `db:"creado_por_id" sanitizer:"id" visible:"false"`
	TituloEvento *string `db:"titulo_evento"`
	FechaInicio *string `db:"fecha_inicio"`
	FechaFin *string `db:"fecha_fin"`
	Tipo *string `db:"tipo"`
	Comentarios *string `db:"comentarios"`
	Lugar *string `db:"lugar"`
}

var Model = base_models.NewModel[EventosStruct]("eventos", "eventos")
