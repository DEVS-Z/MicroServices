// Archivo generado automáticamente para el módulo actividadesUsuario (model)
package actividades_usuario_model

import (
	base_models "main/pkg/base/models"
	actividades_model "main/source/modules/actividades/model"
)

type ActividadesUsuarioStruct struct {
	Id        *int `db:"id" sanitizer:"id" visible:"false"`
	Usuario   *int `db:"usuario"`
	Actividad *int `db:"actividad"`
}

type ActividadesUsuarioStructWithRelacion struct {
	Id        *int
	Usuario   *int
	Actividad actividades_model.ActividadesSanitizer
}

var Model = base_models.NewModel[ActividadesUsuarioStruct]("actividadesUsuario", "actividadesUsuario")
