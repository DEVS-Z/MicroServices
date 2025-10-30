// Archivo generado automáticamente para el módulo Pagos (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type PagosStruct struct {
	PagoId      *int     `db:"pago_id" sanitizer:"id" visible:"false" type: "pk"`
	ClubSubsId  *int     `db:"club_subs_id" sanitizer:"id" visible:"false"`
	TotalPagado *float64 `db:"total_pagado"`
	FechaPago   *string  `db:"fecha_pago"`
}

var Model = base_models.NewModel[PagosStruct]("pagos", "zfut.pagos")
