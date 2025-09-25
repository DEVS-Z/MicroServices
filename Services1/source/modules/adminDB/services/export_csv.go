package handlers

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"main/source/modules/adminDB/schema"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// schema.Credentials debería estar definido como:
// type Credentials struct {
//     Username string `json:"username"`
//     Password string `json:"password"`
// }

type ExportCSVRequest struct {
	Table       string             `json:"table"`
	Credentials schema.Credentials `json:"credentials"`
}

// ExportTableCSVHandler genera un archivo CSV desde una tabla y permite su descarga posterior
func ExportTableCSV(c *gin.Context) {
	var req ExportCSVRequest
	host := os.Getenv("HOST_DB_POSTGRES")
	dbname := os.Getenv("DATABASE_POSTGRES")

	if err := c.BindJSON(&req); err != nil || req.Table == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar el nombre de la tabla y credenciales válidas"})
		return
	}

	if req.Credentials.Username == "" || req.Credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credenciales incompletas"})
		return
	}

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
			host, req.Credentials.Username, req.Credentials.Password, dbname))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a la base de datos"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + req.Table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error consultando la tabla: " + err.Error()})
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo columnas"})
		return
	}

	fileName := fmt.Sprintf("%s_%d.csv", req.Table, time.Now().Unix())
	filePath := filepath.Join("./db/exports", fileName)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando archivo CSV"})
		return
	}

	// Escribir BOM UTF-8 al principio
	file.Write([]byte{0xEF, 0xBB, 0xBF})

	writer := csv.NewWriter(file)

	// Escribe encabezados
	if err := writer.Write(columns); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error escribiendo encabezados CSV"})
		return
	}

	vals := make([]interface{}, len(columns))
	valPtrs := make([]interface{}, len(columns))
	for i := range vals {
		valPtrs[i] = &vals[i]
	}

	for rows.Next() {
		if err := rows.Scan(valPtrs...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo fila"})
			return
		}

		record := make([]string, len(columns))
		for i, val := range vals {
			record[i] = formatCSVValue(val)
		}
		if err := writer.Write(record); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error escribiendo fila CSV"})
			return
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finalizando escritura del CSV"})
		return
	}

	// Asegurarse de que todo fue escrito antes de responder
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "text/csv")
	c.File(filePath)
}

// formatCSVValue convierte valores de diferentes tipos a string
func formatCSVValue(val interface{}) string {
	if val == nil {
		return ""
	}
	switch v := val.(type) {
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	case string:
		return v
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return ""
	}
}

// DownloadCSVHandler permite descargar el CSV generado
func DownloadCSVHandler(c *gin.Context) {
	fileName := c.Param("file")
	filePath := filepath.Join("./db/exports", fileName)

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Archivo no encontrado"})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "text/csv")
	c.File(filePath)
}
