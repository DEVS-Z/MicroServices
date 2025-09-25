package main

import (
	grpcmain "main/cmd/grpc"
	restmain "main/cmd/rest"
	p "main/connection/db/postgres"
	c "main/source/core"
)

func main() {
	// Ejecutar REST en goroutine
	print("Iniciando servicio...")

	go func() {
		p.Init()

		c.Init()
		restmain.RunRest()
	}()

	// Ejecutar gRPC en goroutine
	go func() {
		grpcmain.RunGrpc()
	}()

	// Mantener main corriendo
	select {} // bloquea para que las goroutines sigan ejecutándose
}
