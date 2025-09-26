// Archivo generado automáticamente para el módulo actividadesUsuario (model)
package actividades_usuario_model

import (
	actividades_model "main/source/modules/actividades/model"

	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type ActividadesUsuarioStruct struct {
	Id        *int `db:"id" sanitizer:"id" visible:"false"`
	Usuario   *int `db:"usuario"`
	Actividad *int `db:"actividad"`
}

type ActividadesUsuarioStructWithRelacion struct {
	Id        *int
	Usuario   *string
	Actividad actividades_model.ActividadesStruct
}

var Model = base_models.NewModel[ActividadesUsuarioStruct]("actividadesUsuario", "actividadesUsuario")
