// Archivo generado automáticamente para el módulo permisos (model)
package modulos_model

import (
	base_models "main/pkg/base/models"
)

type ModulosStruct struct {
	Id     *int    `db:"id" sanitizer:"id" visible:"false"`
	Nombre string  `db:"nombre"`
	Status *bool   `db:"status"`
	Icon   *string `db:"icon"`
}

var Model = base_models.NewModel[ModulosStruct]("modulos", "modulos")
