package handlers

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// ResetDBRequest representa la estructura del body esperado
type ResetDBRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	File     string `json:"file"` // opcional: nombre del archivo .sql
}

// ResetDatabase resetea la base de datos usando un archivo .sql
func ResetDatabase(c *gin.Context) {
	var req ResetDBRequest
	host := os.Getenv("HOST_DB_POSTGRES")
	dbname := os.Getenv("DATABASE_POSTGRES")

	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe proporcionar usuario y contraseña"})
		return
	}

	file := req.File
	if file == "" {
		file = "./db/estado_inicial.sql"
	}

	cmd := exec.Command("psql", "-U", req.Username, "-h", host, "-d", dbname, "-f", file)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+req.Password)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reseteando base de datos: " + err.Error(), "details": stderr.String()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Base de datos reseteada a estado inicial"})
}
