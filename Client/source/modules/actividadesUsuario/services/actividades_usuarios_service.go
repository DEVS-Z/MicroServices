package actividades_usuarios_service

import (
	"context"
	"fmt"
	"main/core/connection/services/user_service"
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
	// MODIFICADO: Se usa grpc.NewClient en lugar del grpc.Dial deprecado.
	// La firma es la misma en este caso de uso simple.
	//Mostrando el usuario id
	println("Usuario ID:", userID)

	conn, err := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("no se pudo crear el cliente para el servicio de usuarios: %w", err)
	}
	defer conn.Close()

	client := user_service.NewUserServiceClient(conn)

	// 2️⃣ Hacer la petición ReadOne
	// (El nombre del método era ReadOne en tu proto anterior, lo he corregido aquí)
	resp, err := client.ReadOne(context.Background(), &user_service.UserIdRequest{Id: int32(userID)})
	if err != nil {
		// Se retorna "" en lugar de " " para ser más idiomático en Go.
		return "", fmt.Errorf("error al leer usuario por gRPC: %w", err)
	}

	if resp != nil && resp.User != nil {
		return resp.User.GetPrimerNombre(), nil
	}

	// Si no se encuentra el usuario o la respuesta es vacía, se devuelve un string vacío y sin error.
	return "", nil
}

var Service = base_service.NewService[ActividadesUsuarioService](*actividades_usuario_model.Model)
