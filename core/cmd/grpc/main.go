package grpcmain

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// 1. Crea tu struct para el servidor
// Incrusta UnimplementedUserServiceServer para compatibilidad.
var Server = grpc.NewServer()

func RunGrpc() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al abrir puerto: %v", err)
	}

	log.Println("Servidor gRPC escuchando en :50051")
	if err := Server.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}

}
