package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Column struct {
	Name       string
	Type       string
	Nullable   bool
	IsID       bool
	IsPassword bool
	IsRol      bool
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run sql_to_modules.go <archivo.sql>")
		return
	}

	sqlFile := os.Args[1]

	content, err := os.ReadFile(sqlFile)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
		return
	}

	tables := parseTables(string(content))

	for tableName, columns := range tables {
		err := generateModuleWithStruct(tableName, columns)
		if err != nil {
			fmt.Printf("Error generando módulo para %s: %v\n", tableName, err)
		}
		err = updateCoreIndex(tableName)
		if err != nil {
			fmt.Printf("Error actualizando core/index.go: %v\n", err)
		}
	}

	fmt.Println("✅ Módulos generados correctamente.")
}

func parseTables(sql string) map[string][]Column {
	result := make(map[string][]Column)

	// Regex para capturar CREATE TABLE nombre (...)
	tableRegex := regexp.MustCompile(`(?is)CREATE TABLE\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\((.*?)\);`)
	matches := tableRegex.FindAllStringSubmatch(sql, -1)

	for _, m := range matches {
		tableName := m[1]
		body := m[2]

		cols := parseColumns(body)
		result[tableName] = cols
	}

	return result
}

func parseColumns(body string) []Column {
	lines := strings.Split(body, "\n")
	var cols []Column

	colRegex := regexp.MustCompile(`^\s*([a-zA-Z_][a-zA-Z0-9_]*)\s+([a-zA-Z0-9\(\), ]+)(.*)$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(strings.ToUpper(line), "CONSTRAINT") ||
			strings.HasPrefix(strings.ToUpper(line), "PRIMARY KEY") ||
			strings.HasPrefix(strings.ToUpper(line), "REFERENCES") {
			continue
		}

		m := colRegex.FindStringSubmatch(line)
		if len(m) < 4 {
			continue
		}

		colName := m[1]
		colType := m[2]
		rest := m[3]

		nullable := true
		if strings.Contains(strings.ToUpper(rest), "NOT NULL") {
			nullable = false
		}

		col := Column{
			Name:     colName,
			Type:     colType,
			Nullable: nullable,
			IsID:     strings.EqualFold(colName, "id"),
			// Marcamos manualmente campos sensibles para tags específicos
			IsPassword: strings.EqualFold(colName, "contrasena"),
			IsRol:      strings.EqualFold(colName, "rol"),
		}

		cols = append(cols, col)
	}

	return cols
}

// ==== Helper para conversiones ====

func toPascalCase(s string) string {
	parts := splitIdentifier(s)
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}
	return strings.Join(parts, "")
}

func toLowerCamelCase(s string) string {
	p := toPascalCase(s)
	return strings.ToLower(p[:1]) + p[1:]
}

func toSnakeCaseFromCamel(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

func splitIdentifier(s string) []string {
	s = strings.ReplaceAll(s, "_", " ")
	var parts []string
	runes := []rune(s)
	var current []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && isUpper(runes[i]) && (len(current) > 0 && !isUpper(runes[i-1])) {
			parts = append(parts, string(current))
			current = []rune{}
		}
		if runes[i] == ' ' {
			if len(current) > 0 {
				parts = append(parts, string(current))
				current = []rune{}
			}
		} else {
			current = append(current, runes[i])
		}
	}
	if len(current) > 0 {
		parts = append(parts, string(current))
	}
	return parts
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func sqlTypeToGoType(sqlType string, nullable bool) string {
	// Simplificado, puedes ampliar según tipos que uses
	sqlType = strings.ToUpper(sqlType)

	var goType string

	switch {
	case strings.Contains(sqlType, "INT"), strings.Contains(sqlType, "SERIAL"):
		goType = "int"
	case strings.Contains(sqlType, "VARCHAR"), strings.Contains(sqlType, "TEXT"), strings.Contains(sqlType, "CHAR"):
		goType = "string"
	case strings.Contains(sqlType, "DECIMAL"), strings.Contains(sqlType, "NUMERIC"), strings.Contains(sqlType, "FLOAT"), strings.Contains(sqlType, "DOUBLE"):
		goType = "float64"
	case strings.Contains(sqlType, "BOOLEAN"), strings.Contains(sqlType, "BOOL"):
		goType = "bool"
	case strings.Contains(sqlType, "DATE"), strings.Contains(sqlType, "TIMESTAMP"), strings.Contains(sqlType, "TIME"):
		goType = "time.Time"
	default:
		goType = "string"
	}

	if nullable && goType != "string" {
		goType = "*" + goType
	} else if nullable && goType == "string" {
		// Para strings, opcional dejar como no puntero, pero si quieres puntero, quita este else-if
		// Aquí lo dejo sin puntero para strings nullable (puedes ajustar)
	}

	return goType
}

func generateModuleWithStruct(table string, columns []Column) error {
	structName := toPascalCase(table) + "Struct"
	moduleDir := toLowerCamelCase(table)
	snake := toSnakeCaseFromCamel(table)
	packageName := snake + "_model"

	basePath := filepath.Join("..", "..", "source", "modules", moduleDir)
	modelPath := filepath.Join(basePath, "model")
	os.MkdirAll(modelPath, os.ModePerm)

	indexPath := filepath.Join(basePath, "index.go")
	indexContent := generateIndexGo(moduleDir, packageName, structName, table)
	os.WriteFile(indexPath, []byte(indexContent), 0644)

	modelFileName := snake + "_model.go"
	modelFilePath := filepath.Join(modelPath, modelFileName)
	modelContent := generateModelGoWithFields(packageName, structName, moduleDir, table, columns)
	os.WriteFile(modelFilePath, []byte(modelContent), 0644)

	return nil
}

func generateIndexGo(moduleDir, modelPackage, structName, table string) string {
	return fmt.Sprintf(`package %s

import (
	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	%s "main/source/modules/%s/model"
)

var Service = base_service.NewService[base_service.Default[%s.%s]](*%s.Model)

func Init() {
	print("%s Module Initialized\n")
	InitRoutes()
}

func InitRoutes() {
	var r = router.NewRoute("/%s")
	r.USE(middleware.JWTMiddleware())
	r.GET("/", Service.Read)
	r.POST("/", Service.Insert)
	r.GET("/:id", Service.ReadOne)
	r.PUT("/:id", Service.Update)
	r.DELETE("/:id", Service.Delete)
}
`, moduleDir, modelPackage, moduleDir, modelPackage, structName, modelPackage, toPascalCase(table), table)
}

func generateModelGoWithFields(packageName, structName, lowerCamelName, tableName string, columns []Column) string {
	var fields []string

	for _, c := range columns {
		goType := sqlTypeToGoType(c.Type, c.Nullable)

		fieldName := toPascalCase(c.Name)

		tags := fmt.Sprintf("`db:\"%s\"", strings.ToLower(c.Name))

		if c.IsID {
			tags += ` sanitizer:"id" visible:"false"`
		}
		if c.IsPassword {
			tags += ` sanitizer:"contrasena" visible:"false"`
		}
		if c.IsRol {
			tags += ` sanitizer:"rol" visible:"false"`
		}

		tags += "`"

		fields = append(fields, fmt.Sprintf("\t%s\t%s\t%s", fieldName, goType, tags))
	}

	return fmt.Sprintf(`// Archivo generado automáticamente para el módulo %s (model)
package %s

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
	"time"
)

type %s struct {
%s
}

var Model = base_models.NewModel[%s]("%s", "%s")
`, tableName, packageName, structName, strings.Join(fields, "\n"), structName, lowerCamelName, tableName)
}

func updateCoreIndex(table string) error {
	module := toLowerCamelCase(table)

	indexPath := filepath.Join("..", "..", "source", "core", "index.go")
	contentBytes, err := os.ReadFile(indexPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(contentBytes), "\n")
	importStmt := fmt.Sprintf("\t\"main/source/modules/%s\"", module)
	initCall := fmt.Sprintf("\tNewModule(%s.Init)", module)

	for _, line := range lines {
		if strings.TrimSpace(line) == importStmt || strings.TrimSpace(line) == initCall {
			fmt.Println("⚠️  Ya existe el módulo en core/index.go:", module)
			return nil
		}
	}

	importEnd := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == ")" {
			importEnd = i
			break
		}
	}
	if importEnd != -1 {
		lines = append(lines[:importEnd], append([]string{importStmt}, lines[importEnd:]...)...)
	}

	initIndex := -1
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "func init()") {
			initIndex = i
			break
		}
	}
	if initIndex != -1 {
		for i := initIndex; i < len(lines); i++ {
			if strings.Contains(lines[i], "{") {
				lines = append(lines[:i+1], append([]string{initCall}, lines[i+1:]...)...)
				break
			}
		}
	}

	return os.WriteFile(indexPath, []byte(strings.Join(lines, "\n")), 0644)
}
