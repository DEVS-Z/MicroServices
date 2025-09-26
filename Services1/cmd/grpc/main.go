package grpcmain

import (
	"context"
	"log"
	"net"

	pb "main/connection/services/user_service"
	user_service "main/source/modules/users/services"

	"google.golang.org/grpc"
)

// 1. Crea tu struct para el servidor
// Incrusta UnimplementedUserServiceServer para compatibilidad.
type server struct {
	pb.UnimplementedUserServiceServer
}

// 2. Implementa los métodos de tu servicio
// El nombre y la firma deben coincidir con lo generado por el .proto
func (s *server) ReadOne(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
	log.Printf("Recibida solicitud para el usuario: %v", req.GetId())
	user, err := user_service.Service.ReadOne(int(req.GetId()))
	if err != nil {
		return nil, err
	}

	puser := pb.UserSanitizer{
		Id:              int32(user.ID),
		PrimerNombre:    user.PrimerNombre,
		PrimerApellido:  user.PrimerApellido,
		SegundoApellido: *user.SegundoApellido,
		Correo:          user.Correo,
		Rol:             *user.Rol,
	}

	return &pb.UserResponse{
		User: &puser,
	}, nil
}

func RunGrpc() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al abrir puerto: %v", err)
	}

	s := grpc.NewServer()

	// 3. Registra TU implementación del servidor, no la vacía
	pb.RegisterUserServiceServer(s, &server{})

	log.Println("Servidor gRPC escuchando en :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
