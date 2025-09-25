package handlers

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"path/filepath"

	"main/source/modules/adminDB/schema"

	"github.com/gin-gonic/gin"
)

// BackupFullHandler ejecuta un respaldo completo de la base de datos y retorna el archivo generado
func BackupFull(c *gin.Context) {
	var credentials schema.Credentials

	// Validar el JSON recibido
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El formato de credenciales es inválido"})
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario y contraseña son obligatorios"})
		return
	}

	// Obtener variables de entorno
	host := os.Getenv("HOST_DB_POSTGRES")
	db := os.Getenv("DATABASE_POSTGRES")
	if host == "" || db == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuración del servidor incompleta"})
		return
	}

	// Generar nombre y ruta del archivo
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("respaldo_%s.backup", timestamp)
	outputDir := "./db/respaldos"
	filePath := filepath.Join(outputDir, fileName)

	// Crear directorio si no existe
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el directorio de respaldo"})
		return
	}

	// Ejecutar pg_dump
	cmd := exec.Command("pg_dump", 
		"-U", credentials.Username, 
		"-h", host, 
		"-F", "c",
		"-f", filePath, 
		"-d", db,
	)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+credentials.Password)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al generar el respaldo",
			"details": stderr.String(),
		})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message":     "Respaldo generado exitosamente",
		"backup_file": fileName,
		"path":        filePath,
	})
}


// BackupPartialHandler ejecuta un respaldo parcial de ciertas tablas y retorna el archivo generado
func BackupPartial(c *gin.Context) {
	var params schema.BackupPartialParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	host := os.Getenv("HOST_DB_POSTGRES")
	db := os.Getenv("DATABASE_POSTGRES")

	if params.Credentials.Username == "" || params.Credentials.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credenciales incompletas"})
		return
	}

	if len(params.Tables) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe enviar una lista de tablas"})
		return
	}

	fileName := fmt.Sprintf("respaldo_parcial_%d.backup", time.Now().Unix())
	filePath := "./db/respaldos/" + fileName

	args := []string{"-U", params.Credentials.Username, "-h", host, "-F", "c", "-f", filePath}
	for _, t := range params.Tables {
		args = append(args, "-t", t)
	}
	args = append(args, db)

	cmd := exec.Command("pg_dump", args...)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+params.Credentials.Password)

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error haciendo respaldo parcial: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Respaldo parcial creado correctamente",
		"backup_file":      fileName,
		"tablas_incluidas": params.Tables,
	})
}

func ListBackups(c *gin.Context) {
	backupDir := "./db/respaldos"
	files, err := os.ReadDir(backupDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo el directorio de respaldos"})
		return
	}

	var backups []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".backup") {
			backups = append(backups, file.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"backups": backups,
	})
}

// ListTables lista las tablas disponibles en el esquema 'public'
func ListTables(c *gin.Context) {
	var creds schema.Credentials
	host := os.Getenv("HOST_DB_POSTGRES")
	dbname := os.Getenv("DATABASE_POSTGRES")

	if err := c.ShouldBindJSON(&creds); err != nil || creds.Username == "" || creds.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe proporcionar usuario y contraseña"})
		return
	}

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, creds.Username, creds.Password, dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error conectando a la base de datos"})
		return
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
		AND table_type = 'BASE TABLE'
		ORDER BY table_name
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo las tablas"})
		return
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error leyendo resultado"})
			return
		}
		tables = append(tables, table)
	}

	c.JSON(http.StatusOK, gin.H{"tables": tables})
}

