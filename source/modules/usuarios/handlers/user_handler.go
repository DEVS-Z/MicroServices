package handlers

import (
	"context"
	"log"
	grpcmain "main/core/cmd/grpc"
	pb "main/core/connection/services/user_service"
	usuarios_service "main/source/modules/usuarios/service"
)

type Server struct {
	usuarios_service.ServiceType
	pb.UnimplementedUserServiceServer
}

func InitGrpcServer(s usuarios_service.ServiceType) {
	pb.RegisterUserServiceServer(grpcmain.Server, &Server{
		ServiceType: s,
	})
}

// 2. Implementa los métodos de tu servicio
// El nombre y la firma deben coincidir con lo generado por el .proto
func (s *Server) ReadOne(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
	log.Printf("Recibida solicitud para el usuario: %v", req.GetUserId())
	user, err := s.Service.ReadOne(int(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	puser := pb.UserSanitizer{
		UserId:        int32(*user.UserId),
		Nombre:        *user.Nombre,
		Correo:        *user.Correo,
		FechaRegistro: *user.FechaRegistro,
		Estado:        *user.Estado,
		RolId:         int32(*user.RolId),
	}

	return &pb.UserResponse{
		User: &puser,
	}, nil
}
