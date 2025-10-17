package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
	ORM "github.com/miqueaz/FrameGo/pkg/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB = PostgreSQL()

func PostgreSQL() *sql.DB {
	var err error
	godotenv.Load("./core/config/.env")
	connection := ORM.Connection{
		Host:     os.Getenv("HOST_DB_POSTGRES"),
		Port:     os.Getenv("PORT_DB_POSTGRES"),
		User:     os.Getenv("USER_DB_POSTGRES"),
		Password: os.Getenv("PASSWORD_DB_POSTGRES"),
		Database: os.Getenv("DATABASE_POSTGRES"),
		SSLMode:  os.Getenv("SSLMODE_DB_POSTGRES"),
	}
	println("Conectando a la base de datos PostgreSQL...", connection.Host)
	DB, err := ORM.InitPostgres(connection)
	if err != nil || DB == nil {
		log.Fatalf("Error inicializando PostgreSQL: %v", err)
	}
	base_models.SetDB(sqlx.NewDb(DB, "postgres"))
	if DB == nil {
		log.Fatal("Error: La conexión a la base de datos PostgreSQL no se ha inicializado." + err.Error())
	}

	println("Conexión a PostgreSQL exitosa")
	ListTables(DB)
	return DB

}

func ListTables(db *sql.DB) {
	query := `
        SELECT table_schema, table_name
        FROM information_schema.tables
        WHERE table_type='BASE TABLE'
        AND table_schema NOT IN ('pg_catalog', 'information_schema')
        ORDER BY table_schema, table_name;
    `

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error listando tablas: %v", err)
	}
	defer rows.Close()

	fmt.Println("📋 Tablas en la base de datos:")
	for rows.Next() {
		var schema, name string
		err := rows.Scan(&schema, &name)
		if err != nil {
			log.Fatalf("Error leyendo fila: %v", err)
		}
		fmt.Printf("- %s.%s\n", schema, name)
	}
}
