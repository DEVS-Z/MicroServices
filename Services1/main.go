package main

import (
	"fmt"
	grpcmain "main/cmd/grpc"
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

	// Ejecutar gRPC en goroutine
	go func() {
		grpcmain.RunGrpc()
	}()

	// Mantener main corriendo
	select {} // bloquea para que las goroutines sigan ejecutándose
}
