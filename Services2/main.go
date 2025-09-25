package main

import (
	"fmt"
	restmain "main/cmd/rest"
	"main/config"
	_ "main/connection/db/postgres"
	"main/source/core"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func main() {
	// Ejecutar REST en goroutine
	fmt.Println("Iniciando servicio...")
	config.Execute()
	core.Init()
	modules.Execute()

	go func() {
		restmain.RunRest()
	}()

	// Mantener main corriendo
	select {} // bloquea para que las goroutines sigan ejecutándose
}
