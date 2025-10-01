package user_handlers

import (
	"context"
	"log"
	pb "main/source/modules/users/proto/user_service"
	user_service "main/source/modules/users/services"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

// 2. Implementa los métodos de tu servicio
// El nombre y la firma deben coincidir con lo generado por el .proto
func (s *Server) ReadOne(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
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
