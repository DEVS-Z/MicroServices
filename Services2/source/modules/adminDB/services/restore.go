package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type RestoreBackupRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	BackupFile string `json:"backup_file"`
}

// RestoreBackup permite restaurar un respaldo de la base de datos usando pg_restore
func RestoreBackup(c *gin.Context) {
	var req RestoreBackupRequest

	// Validar JSON de entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El formato de la petición es inválido"})
		return
	}

	if req.Username == "" || req.Password == "" || req.BackupFile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe proporcionar usuario, contraseña y nombre del archivo de respaldo"})
		return
	}

	ok, msg := checkRestorePreconditions(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	host := os.Getenv("HOST_DB_POSTGRES")
	dbname := os.Getenv("DATABASE_POSTGRES")

	backupDir := "./db/respaldos"
	filePath := filepath.Join(backupDir, req.BackupFile)

	// Restaurar con pg_restore
	cmd := exec.Command("pg_restore",
		"-U", req.Username,
		"-h", host,
		"-d", dbname,
		"-c",         // limpiar objetos antes de crear
		"--no-owner", // evita errores si el dueño no existe
		filePath,
	)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+req.Password)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al restaurar el respaldo",
			"details": stderr.String(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Respaldo restaurado correctamente",
		"restored_db": dbname,
		"file":        req.BackupFile,
	})
}

func checkRestorePreconditions(req RestoreBackupRequest) (bool, string) {
	// 1. Validar variables de entorno
	host := os.Getenv("HOST_DB_POSTGRES")
	dbname := os.Getenv("DATABASE_POSTGRES")

	if host == "" || dbname == "" {
		return false, "Faltan variables de entorno: HOST_DB_POSTGRES o DATABASE_POSTGRES"
	}

	// 2. Validar nombre seguro de archivo (sin rutas maliciosas)
	if strings.Contains(req.BackupFile, "..") || strings.ContainsAny(req.BackupFile, `/\`) {
		return false, "Nombre de archivo inválido por razones de seguridad"
	}

	// 3. Validar existencia del archivo
	backupDir := "./db/respaldos"
	filePath := filepath.Join(backupDir, req.BackupFile)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false, fmt.Sprintf("El archivo '%s' no existe", req.BackupFile)
	}

	// 4. Validar que pg_restore puede leer el archivo
	cmd := exec.Command("pg_restore", "-l", filePath)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+req.Password, "PGUSER="+req.Username)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return false, "No se pudo leer el archivo de respaldo. Puede estar corrupto o ser incompatible."
	}

	// 5. Validar conexión a la base de datos
	testConn := exec.Command("psql", "-U", req.Username, "-h", host, "-d", dbname, "-c", "\\q")
	testConn.Env = append(os.Environ(), "PGPASSWORD="+req.Password)
	if err := testConn.Run(); err != nil {
		return false, "No se pudo conectar a la base de datos con las credenciales proporcionadas"
	}

	return true, ""
}
