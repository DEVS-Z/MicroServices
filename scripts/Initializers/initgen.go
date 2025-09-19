package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run initgen.go <nombre_del_modulo>")
		os.Exit(1)
	}

	module := strings.ToLower(os.Args[1])

	err := generateModule(module)
	if err != nil {
		fmt.Printf("Error al generar módulo: %v\n", err)
		os.Exit(1)
	}

	err = updateCoreIndex(module)
	if err != nil {
		fmt.Printf("Error al actualizar core/index.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Módulo '%s' generado y registrado en core/index.go correctamente.\n", module)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func generateModule(module string) error {
	basePath := filepath.Join("..", "..", "source", "modules", module)
	modelPath := filepath.Join(basePath, "model")

	err := os.MkdirAll(modelPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Crear index.go
	indexContent := generateIndexGo(module)
	indexFile := filepath.Join(basePath, "index.go")
	err = os.WriteFile(indexFile, []byte(indexContent), 0644)
	if err != nil {
		return err
	}

	// Crear model.go
	modelContent := generateModelGo(module)
	modelFile := filepath.Join(modelPath, module+"_model.go")
	err = os.WriteFile(modelFile, []byte(modelContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateIndexGo(module string) string {
	mod := strings.ToLower(module)
	structName := strings.Title(mod)
	modelImport := fmt.Sprintf(`%s_model "main/source/modules/%s/model"`, mod, mod)

	return fmt.Sprintf(`package %s

import (
	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
	middleware "main/source/helpers/middlewares"
	"main/source/helpers/router"
	%s
)

var Service = base_service.NewService[base_service.Default[%s_model.%sStruct]](*%s_model.Model)

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
`, mod, modelImport, mod, structName, mod, mod, mod)
}

func generateModelGo(module string) string {
	mod := strings.ToLower(module)
	structName := strings.Title(mod) + "Struct"

	return fmt.Sprintf(`// Archivo generado automáticamente para el módulo %s (model)
package %s_model

import (
	base_models "github.com/miqueaz/FrameGo/pkg/base/models"
)

type %s struct {
}

var Model = base_models.NewModel[%s]("%s", "%s")
`, mod, mod, structName, structName, mod, mod)
}

func updateCoreIndex(module string) error {
	indexPath := filepath.Join("..", "..", "source", "modules", "core", "index.go")
	contentBytes, err := os.ReadFile(indexPath)
	if err != nil {
		return err
	}

	content := string(contentBytes)
	lines := strings.Split(content, "\n")

	importStmt := fmt.Sprintf("\t\"main/source/modules/%s\"", module)
	newModuleLine := fmt.Sprintf("\tNewModule(%s.Init)", module)

	// Evitar duplicados
	for _, line := range lines {
		if strings.TrimSpace(line) == importStmt || strings.TrimSpace(line) == newModuleLine {
			fmt.Println("ℹ️  El módulo ya está registrado en core/index.go")
			return nil
		}
	}

	// Insertar import
	importIndex := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == ")" {
			importIndex = i
			break
		}
	}
	if importIndex == -1 {
		return fmt.Errorf("no se encontró el bloque de imports en core/index.go")
	}
	lines = append(lines[:importIndex], append([]string{importStmt}, lines[importIndex:]...)...)

	// Insertar NewModule
	initIndex := -1
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "func init()") {
			initIndex = i
			break
		}
	}
	if initIndex == -1 {
		return fmt.Errorf("no se encontró la función init en core/index.go")
	}

	for i := initIndex; i < len(lines); i++ {
		if strings.Contains(lines[i], "{") {
			lines = append(lines[:i+1], append([]string{newModuleLine}, lines[i+1:]...)...)
			break
		}
	}

	newContent := strings.Join(lines, "\n")
	return os.WriteFile(indexPath, []byte(newContent), 0644)
}
