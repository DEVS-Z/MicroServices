package actividades_usuarios_service

import (
	"context"
	"fmt"
	"main/connection/services/user_service"
	actividades_services "main/source/modules/actividades/services"
	actividades_usuario_model "main/source/modules/actividadesUsuario/model"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ActividadesUsuarioService struct {
	base_service.Service[actividades_usuario_model.ActividadesUsuarioStruct]
}

func (s *ActividadesUsuarioService) Read(filters map[string]interface{}) ([]actividades_usuario_model.ActividadesUsuarioStructWithRelacion, error) {

	var result []actividades_usuario_model.ActividadesUsuarioStructWithRelacion
	data, err := s.Service.Read(filters)
	if err != nil {
		return nil, err
	}

	for _, item := range data {
		var newItem actividades_usuario_model.ActividadesUsuarioStructWithRelacion
		actividad, err := actividades_services.Service.ReadOne(*item.Actividad)
		if err != nil {
			return nil, err
		}

		nombreUsuario, err := GetUserName(*item.Usuario)
		if err != nil {
			return nil, err
		}

		newItem.Id = item.Id
		newItem.Usuario = &nombreUsuario
		newItem.Actividad = actividad
		result = append(result, newItem)
	}

	return result, nil
}

func GetUserName(userID int) (string, error) {
	// 1️⃣ Conectar al microservicio de usuarios por gRPC
	conn, err := grpc.Dial("service1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return " ", fmt.Errorf("no se pudo conectar a usuarios: %w", err)
	}
	defer conn.Close()

	client := user_service.NewUserServiceClient(conn)

	// 2️⃣ Hacer la petición Read (todos los usuarios)
	resp, err := client.ReadOne(context.Background(), &user_service.UserIdRequest{Id: int32(userID)})
	if err != nil {
		return " ", fmt.Errorf("error al leer usuarios por gRPC: %w", err)
	}

	if resp != nil && resp.User != nil {
		return resp.User.GetPrimerNombre(), nil
	}

	return "", nil
}

var Service = base_service.NewService[ActividadesUsuarioService](*actividades_usuario_model.Model)
