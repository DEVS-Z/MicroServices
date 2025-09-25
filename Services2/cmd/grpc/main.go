package grpcmain

import (
	"context"
	"log"
	"time"

	pb "main/connection/services/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGrpc() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Crear la conexión usando NewClient
	conn, err := grpc.DialContext(
		ctx,
		"service1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()), // reemplaza WithInsecure
	)
	if err != nil {
		log.Fatalf("no se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Llamada a Read
	resp, err := client.Read(ctx, &pb.UserFilterRequest{})
	if err != nil {
		log.Fatalf("error leyendo usuarios: %v", err)
	}

	for _, u := range resp.Users {
		log.Printf("ID: %d, Nombre: %s, Rol: %s", u.Id, u.PrimerNombre, u.Rol)
	}
}
