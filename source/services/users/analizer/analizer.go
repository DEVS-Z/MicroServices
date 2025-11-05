package analizer

import (
	"context"
	"fmt"
	"main/core/connection/services/user_service/analizer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetUserName(crudo map[string]any) (string, error) {

	fmt.Print("Conectando...\n")
	conn, err := grpc.NewClient("172.18.4.28:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("no se pudo crear el cliente para el servicio de usuarios: %w", err)
	}
	defer conn.Close()
	client := analizer.NewUserServiceClient(conn)

	// Convert map to structpb.Struct
	structData, err := structpb.NewStruct(crudo)
	if err != nil {
		return "", fmt.Errorf("error converting map to struct: %w", err)
	}

	payload, err := anypb.New(structData)
	if err != nil {
		return "", fmt.Errorf("error creating Any payload: %w", err)
	}

	userData := &analizer.UserData{
		Payload: payload,
	}

	// 2️⃣ Hacer la petición InsertUserData
	resp, err := client.InsertUserData(context.Background(), userData)
	if err != nil {
		// Se retorna "" en lugar de " " para ser más idiomático en Go.
		return "", fmt.Errorf("error al leer usuario por gRPC: %w", err)
	}

	if resp != nil && resp.Ok != false {
		return resp.GetMessage(), nil
	}

	fmt.Print("No se pudo insertar los datos del usuario\n")
	// Si no se encuentra el usuario o la respuesta es vacía, se devuelve un string vacío y sin error.
	return "", nil
}
