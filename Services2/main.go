package main

import (
	restmain "main/cmd/rest"
	p "main/connection/db/postgres"
	c "main/source/core"
)

func main() {
	// Ejecutar REST en goroutine
	p.Init()
	c.Init()
	go func() {
		restmain.RunRest()
	}()

	// Mantener main corriendo
	select {} // bloquea para que las goroutines sigan ejecutándose
}
