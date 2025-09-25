package util

import (
	"context"
	"fmt"
	pb "main/connection/services/user_service" // structs proto
	roles_model "main/source/modules/roles/model"
	roles_services "main/source/modules/roles/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UsersByRole() ([]roles_model.UsersInRoles, error) {
	// 1️⃣ Conectar al microservicio de usuarios por gRPC
	conn, err := grpc.Dial("service1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar a usuarios: %w", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// 2️⃣ Hacer la petición Read (todos los usuarios)
	resp, err := client.Read(context.Background(), &pb.UserFilterRequest{})
	if err != nil {
		return nil, fmt.Errorf("error al leer usuarios por gRPC: %w", err)
	}

	// 3️⃣ Obtener roles internos
	rolesArr, err := roles_services.Service.Read(nil)
	if err != nil || len(rolesArr) == 0 {
		return nil, fmt.Errorf("error al leer roles: %w", err)
	}

	// 4️⃣ Mapear usuarios a cada rol
	dataSanitizer := make([]roles_model.UsersInRoles, 0, len(rolesArr))
	for _, role := range rolesArr {
		var users []*pb.UserSanitizer
		for _, user := range resp.Users { // resp.Users ya es []*UserSanitizer
			if user.Rol == role.Nombre {
				users = append(users, user)
			}
		}

		roleName := role.Nombre
		dataSanitizer = append(dataSanitizer, roles_model.UsersInRoles{
			Role:  &roleName,
			Users: users, // aquí puedes cambiar Users a []*pb.UserSanitizer o mapear si quieres tu modelo interno
		})
	}

	return dataSanitizer, nil
}
