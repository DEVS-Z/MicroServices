package auth

import (
	"errors"
	key "main/core/security/token"
	user "main/source/modules/usuarios"
	"time"

	"github.com/miqueaz/FrameGo/pkg/crypto"
)

type AuthService struct {
	email    string
	password string
}

func SignIn(crudo map[string]any) (string, error) {

	//Transformar el body a AuthService
	body := AuthService{
		email:    crudo["email"].(string),
		password: crudo["password"].(string),
	}

	users, err := user.Service.Read(map[string]any{"correo": body.email})

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

	if *user.Password != body.password {
		return "", errors.New("invalid password")
	}

	token, err := key.GenerateJWT(*user.Nombre, *user.RolId, *user.UserId)
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

	estado := "activo"
	rol := 5
	fechaRegistro := string(time.Now().Format("2006-01-02"))

	_, err = user.Service.Insert(user.Model{
		Nombre:        &username,
		Correo:        &email,
		Password:      &hashedPassword,
		FechaRegistro: &fechaRegistro,
		Estado:        &estado,
		RolId:         &rol,
	})

	if err != nil {
		return "", err
	}

	return "Registro Exitoso", nil
}
