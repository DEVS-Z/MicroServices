package grpcmain

import (
	"log"
	"net"

	pb "main/connection/services/user_service"

	"google.golang.org/grpc"
)

func RunGrpc() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al abrir puerto: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, pb.UnimplementedUserServiceServer{})

	log.Println("Servidor gRPC escuchando en :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
