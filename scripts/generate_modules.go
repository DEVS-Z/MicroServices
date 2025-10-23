package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const modulesPath = "../source/modules"
const coreFile = "../source/core/index.go"
const sqlFile = "../schema.sql"

var indexTemplate = `package {name}

import (
	jwt_middleware "main/source/helpers/middlewares/jwt"
	"main/source/helpers/router"
	{name}_model "main/source/modules/{name}/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

var Service = base_service.NewService[base_service.Default[{name}_model.{structName}Struct]](*{name}_model.Model)

func Init() {
	print("{structName} Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/{name}")
	r.USE(jwt_middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
`

var modelTemplate = `// Archivo generado automáticamente para el módulo {structName} (model)
package model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type {structName}Struct struct {
{fields}
}

var Model = base_models.NewModel[{structName}Struct]("{table}", "{table}")
`

func main() {
	content, err := os.ReadFile(sqlFile)
	if err != nil {
		fmt.Println("❌ Error al abrir el archivo SQL:", err)
		return
	}

	tables := parseTables(string(content))

	if len(tables) == 0 {
		fmt.Println("❌ No se encontraron tablas.")
		return
	}

	for _, t := range tables {
		createModule(t)
		updateCoreFile(t.Name)
	}

	fmt.Println("✅ Módulos generados correctamente.")
}

type Table struct {
	Name    string
	Columns []Column
}

type Column struct {
	Name string
	Type string
}

func parseTables(sql string) []Table {
	reTable := regexp.MustCompile(`(?i)CREATE TABLE\s+zfut\.([a-zA-Z_]+)\s*\(([\s\S]*?)\);`)
	reColumn := regexp.MustCompile(`\s*([a-zA-Z_]+)\s+([a-zA-Z0-9\(\)]+)`)

	matches := reTable.FindAllStringSubmatch(sql, -1)
	var tables []Table

	for _, match := range matches {
		tableName := strings.TrimSpace(match[1])
		columnDefs := match[2]
		lines := strings.Split(columnDefs, "\n")

		var columns []Column
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "--") || strings.HasPrefix(line, "CONSTRAINT") {
				continue
			}
			colMatch := reColumn.FindStringSubmatch(line)
			if len(colMatch) > 2 {
				colName := strings.TrimSpace(colMatch[1])
				colType := strings.ToUpper(colMatch[2])
				columns = append(columns, Column{Name: colName, Type: colType})
			}
		}

		tables = append(tables, Table{Name: tableName, Columns: columns})
	}
	return tables
}

func createModule(t Table) {
	structName := strings.Title(t.Name)
	moduleDir := filepath.Join(modulesPath, t.Name)
	modelDir := filepath.Join(moduleDir, "model")
	os.MkdirAll(modelDir, os.ModePerm)

	// Crear index.go
	indexPath := filepath.Join(moduleDir, "index.go")
	indexContent := strings.ReplaceAll(indexTemplate, "{name}", t.Name)
	indexContent = strings.ReplaceAll(indexContent, "{structName}", structName)
	os.WriteFile(indexPath, []byte(indexContent), 0644)

	// Crear model
	fields := buildFields(t.Columns)
	modelPath := filepath.Join(modelDir, fmt.Sprintf("%s_model.go", t.Name))
	modelContent := strings.ReplaceAll(modelTemplate, "{structName}", structName)
	modelContent = strings.ReplaceAll(modelContent, "{table}", t.Name)
	modelContent = strings.ReplaceAll(modelContent, "{fields}", fields)
	os.WriteFile(modelPath, []byte(modelContent), 0644)

	fmt.Printf("📦 Módulo creado: %s\n", t.Name)
}

func buildFields(columns []Column) string {
	var lines []string
	for _, c := range columns {
		goType := sqlToGoType(c.Type)
		fieldName := toCamel(c.Name)

		tag := fmt.Sprintf("`db:\"%s\"`", c.Name)

		// Caso especial para id
		if strings.Contains(strings.ToLower(c.Name), "id") && (strings.HasSuffix(strings.ToLower(c.Name), "_id") || c.Name == "id") {
			tag = fmt.Sprintf("`db:\"%s\" sanitizer:\"id\" visible:\"false\"`", c.Name)
		}

		line := fmt.Sprintf("\t%s %s %s", fieldName, goType, tag)
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func sqlToGoType(sqlType string) string {
	switch {
	case strings.Contains(sqlType, "BIGINT"), strings.Contains(sqlType, "INT"):
		return "*int"
	case strings.Contains(sqlType, "DECIMAL"), strings.Contains(sqlType, "FLOAT"):
		return "*float64"
	case strings.Contains(sqlType, "BIT"):
		return "*bool"
	case strings.Contains(sqlType, "DATETIME"), strings.Contains(sqlType, "DATE"):
		return "*string"
	case strings.Contains(sqlType, "NVARCHAR"), strings.Contains(sqlType, "VARCHAR"), strings.Contains(sqlType, "TEXT"):
		return "*string"
	default:
		return "*string"
	}
}

func toCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func updateCoreFile(moduleName string) {
	contentBytes, err := os.ReadFile(coreFile)
	if err != nil {
		fmt.Println("⚠️ No se pudo abrir core/index.go:", err)
		return
	}
	content := string(contentBytes)

	// Imports del módulo y del modelo
	importModule := fmt.Sprintf("\t\"main/source/modules/%s\"", moduleName)
	importModel := fmt.Sprintf("\t\"main/source/modules/%s/model\"", moduleName)

	// Línea Init
	initLine := fmt.Sprintf("\tmodules.NewModule(%s.Init)", moduleName)

	// Insertar imports
	if !strings.Contains(content, importModule) {
		content = insertImport(content, importModule)
	}
	if !strings.Contains(content, importModel) {
		content = insertImport(content, importModel)
	}

	// Insertar Init
	if !strings.Contains(content, initLine) {
		content = insertInit(content, initLine)
	}

	os.WriteFile(coreFile, []byte(content), 0644)
	fmt.Printf("🔧 core/index.go actualizado con %s y su modelo\n", moduleName)
}

func insertImport(content, line string) string {
	re := regexp.MustCompile(`import\s*\(\s*([\s\S]*?)\)`)
	return re.ReplaceAllString(content, fmt.Sprintf("import (\n$1\n%s\n)", line))
}

func insertInit(content, line string) string {
	re := regexp.MustCompile(`func Init\(\) \{\s*([\s\S]*?)\}`)
	return re.ReplaceAllString(content, fmt.Sprintf("func Init() {\n$1\n%s\n}", line))
}
