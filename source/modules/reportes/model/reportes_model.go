// Archivo generado automáticamente para el módulo Reportes (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type ReportesStruct struct {
	ReporteId     *int    `db:"reporte_id" sanitizer:"id" visible:"false" type: "pk"`
	JugadorId     *int    `db:"jugador_id" sanitizer:"id" visible:"false"`
	FechaRegistro *string `db:"fecha_registro"`
	Tipo          *string `db:"tipo"`
	Comentarios   *string `db:"comentarios"`
}

var Model = base_models.NewModel[ReportesStruct]("reportes", "zfut.reportes")
