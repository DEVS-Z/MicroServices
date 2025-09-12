package admindb

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	handlers "main/source/modules/adminDB/services"
)

func Init() {
	print("Admin DB Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/db/admin")
	r.USE(jwt_middleware.JWTMiddleware())
	r.POST("/backup/full", handlers.BackupFull)
	r.POST("/backup/partial", handlers.BackupPartial)
	r.POST("/export/csv", handlers.ExportTableCSV)
	r.POST("/reset", handlers.ResetDatabase)
	r.POST("/restore", handlers.RestoreBackup)
	r.POST("/backup/list", handlers.ListBackups)
	r.POST("/tables", handlers.ListTables)
}
