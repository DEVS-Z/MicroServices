package usuarios

import (
	usuarios_model "main/source/modules/usuarios/model"
	"main/source/modules/usuarios/routes"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

type Model = usuarios_model.UsuariosStruct
type ServiceType = base_service.Default[Model]

var Service = base_service.NewService[ServiceType](*usuarios_model.Model)

func Init() {
	print("Usuarios Module Initialized\n")
	routes.InitRoutes(Service)
}
