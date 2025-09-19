package postgres

import (
	"database/sql"
	"log"
	"os"

	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
	ORM "github.com/miqueaz/FrameGo/pkg/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	godotenv.Load("./config/.env")
	connection := ORM.Connection{
		Host:     os.Getenv("HOST_DB_POSTGRES"),
		Port:     os.Getenv("PORT_DB_POSTGRES"),
		User:     os.Getenv("USER_DB_POSTGRES"),
		Password: os.Getenv("PASSWORD_DB_POSTGRES"),
		Database: os.Getenv("DATABASE_POSTGRES"),
		SSLMode:  os.Getenv("SSLMODE_DB_POSTGRES"),
	}
	println("Conectando a la base de datos PostgreSQL...", connection.Host)
	DB, err = ORM.InitPostgres(connection)
	base_models.SetGlobalDB(DB)
	if DB == nil {
		log.Fatal("Error: La conexión a la base de datos PostgreSQL no se ha inicializado." + err.Error())
	}
}
