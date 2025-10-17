package main

import (
	"fmt"
	grpcmain "main/core/cmd/grpc"
	restmain "main/core/cmd/rest"
	_ "main/core/connection/db/mysql"
	"main/source/core"

	modules "github.com/miqueaz/FrameGo/pkg/base/core"
)

func main() {
	// Ejecutar REST en goroutine
	fmt.Println("Iniciando servicio...")
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
