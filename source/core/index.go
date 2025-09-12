package modules

import (
	"main/source/helpers/auth"
	"main/source/modules/actividades"
	"main/source/modules/actividadesUsuario"
	adminDB "main/source/modules/adminDB"
	"main/source/modules/alertaMedidores"
	"main/source/modules/alertas"
	"main/source/modules/archivos"
	"main/source/modules/areasIndustriales"
	"main/source/modules/bitacoras"
	"main/source/modules/estadoActividad"
	"main/source/modules/fuenteLectura"
	"main/source/modules/industria"
	"main/source/modules/lecturas"
	"main/source/modules/lineasAgua"
	"main/source/modules/medidores"
	"main/source/modules/mensajes"
	"main/source/modules/metas"
	"main/source/modules/modulos"
	modulesRol "main/source/modules/modulosRol"
	"main/source/modules/notificaciones"
	"main/source/modules/reportes"
	"main/source/modules/roles"
	"main/source/modules/tipoActividad"
	"main/source/modules/tipoAlerta"
	"main/source/modules/tipoArchivo"
	"main/source/modules/tipoBitacora"
	"main/source/modules/tipoMedidor"
	"main/source/modules/unidadMedida"
	"main/source/modules/users"
)

func init() {
	NewModule(actividades.Init)
	NewModule(estadoActividad.Init)
	NewModule(modulesRol.Init)
	NewModule(tipoAlerta.Init)
	NewModule(tipoArchivo.Init)
	NewModule(alertaMedidores.Init)
	NewModule(alertas.Init)
	NewModule(reportes.Init)
	NewModule(medidores.Init)
	NewModule(tipoMedidor.Init)
	NewModule(modulos.Init)
	NewModule(archivos.Init)
	NewModule(metas.Init)
	NewModule(lineasAgua.Init)
	NewModule(lecturas.Init)
	NewModule(unidadMedida.Init)
	NewModule(fuenteLectura.Init)
	NewModule(tipoActividad.Init)
	NewModule(roles.Init)
	NewModule(areasIndustriales.Init)
	NewModule(industria.Init)
	NewModule(bitacoras.Init)
	NewModule(tipoBitacora.Init)
	NewModule(actividadesUsuario.Init)
	NewModule(users.Init)
	NewModule(adminDB.Init)
	NewModule(notificaciones.Init)
	NewModule(mensajes.Init)
	auth.AuthRouter()

}
