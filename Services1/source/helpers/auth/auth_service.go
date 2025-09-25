package auth

import (
	"errors"
	key "main/security/token"
	user_model "main/source/modules/users/models"
	user_service "main/source/modules/users/services"

	"github.com/miqueaz/FrameGo/pkg/crypto"
)

type AuthService struct {
	username string
	email    string
	password string
}

func SignIn(crudo map[string]any) (string, error) {

	//Transformar el body a AuthService
	body := AuthService{
		username: crudo["username"].(string),
		email:    crudo["email"].(string),
		password: crudo["password"].(string),
	}

	users, err := user_service.Service.Service.Read(map[string]any{"Matricula": body.username})
	if len(users) <= 0 {
		return "", errors.New("user not found")
	}
	user := users[0]
	if err != nil {
		return "", err
	}

	// if err := crypto.CheckPassword(user.Contrasena, body.password); err != nil {
	// 	return "", errors.New("invalid password")
	// }

	if user.Contrasena != body.password {
		return "", errors.New("invalid password")
	}

	token, err := key.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func SignUp(username string, email string, password string) (string, error) {

	hashedPassword, err := crypto.EncryptPassword(password)
	if err != nil {
		return "", err
	}

	user, err := user_service.Service.Insert(user_model.UserStruct{
		PrimerNombre:    username,
		SegundoNombre:   nil,
		PrimerApellido:  username,
		SegundoApellido: nil,
		Matricula:       username,
		Correo:          email,
		Contrasena:      hashedPassword,
		Rol:             1, // Assuming 1 is the default role for new users
	})

	if err != nil {
		return "", err
	}

	return user.Matricula, nil
}
