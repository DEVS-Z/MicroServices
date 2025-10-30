package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
	ORM "github.com/miqueaz/FrameGo/pkg/sql"

	"github.com/joho/godotenv"
)

var DB = MySQL()

func MySQL() *sql.DB {
	var err error
	godotenv.Load("./core/config/.env")
	connection := ORM.Connection{
		Host:     os.Getenv("HOST_DB_MYSQL"),
		Port:     os.Getenv("PORT_DB_MYSQL"),
		User:     os.Getenv("USER_DB_MYSQL"),
		Password: os.Getenv("PASSWORD_DB_MYSQL"),
		Database: os.Getenv("DATABASE_MYSQL"),
		SSLMode:  os.Getenv("SSLMODE_DB_MYSQL"),
	}
	println("Conectando a la base de datos MySQL...", connection.Host)
	DB, err := ORM.InitMySQL(connection)
	if err != nil || DB == nil {
		log.Printf("Error inicializando MySQL: %v", err)
		return nil
	}
	base_models.SetDB(sqlx.NewDb(DB, "mysql"))
	if DB == nil {
		log.Print("Error: La conexión a la base de datos MySQL no se ha inicializado." + err.Error())
		return nil
	}

	println("Conexión a MySQL exitosa")
	return DB

}
