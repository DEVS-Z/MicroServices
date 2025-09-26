// Archivo generado automáticamente para el módulo actividades (model)
package actividades_model

import (
	"time"

	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type ActividadesStruct struct {
	Id          *int       `db:"id" sanitizer:"id" visible:"false"`
	Titulo      string     `db:"titulo"`
	Descripcion string     `db:"descripcion"`
	FechaInicio *time.Time `db:"fechainicio"`
	FechaFin    *time.Time `db:"fechafin"`
	Tipo        *int       `db:"tipo"`
	Estado      *int       `db:"estado"`
}

var Model = base_models.NewModel[ActividadesStruct]("actividades", "actividades")
