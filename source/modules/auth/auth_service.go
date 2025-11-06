package auth

import (
	"errors"
	"fmt"
	key "main/core/security/token"
	user "main/source/modules/usuarios"
	"regexp"
	"time"

	"github.com/go-sql-driver/mysql"
)

type AuthService struct {
	email    string
	password string
}

var getStr = func(m map[string]interface{}, k string) string {
	if v, ok := m[k].(string); ok {
		return v
	}
	return ""
}

func SignIn(crudo map[string]any) (string, error) {

	//Transformar el body a AuthService
	body := AuthService{
		email:    getStr(crudo, "email"),
		password: getStr(crudo, "password"),
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

func SignUp(crudo map[string]any) (string, error) {

	nombre := getStr(crudo, "nombre")
	correo := getStr(crudo, "email")
	password := getStr(crudo, "password")

	if nombre == "" || password == "" || correo == "" {
		return "", errors.New("missing fields")
	}

	// hashedPassword, err := crypto.EncryptPassword(password)
	// if err != nil {
	// 	return "", err
	// }

	estado := "activo"
	rol := 5
	fechaRegistro := string(time.Now().Format("2006-01-02"))

	_, err := user.Service.Insert(user.Model{
		Nombre:        &nombre,
		Correo:        &correo,
		Password:      &password,
		FechaRegistro: &fechaRegistro,
		Estado:        &estado,
		RolId:         &rol,
	})

	if err != nil {
		return "", HandleDBError(err)
	}

	return "Registro Exitoso", nil
}

// HandleDBError detecta duplicados y devuelve el campo afectado
func HandleDBError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == 1062 { // Duplicate entry
			// Usamos regex para extraer el campo de la clave
			re := regexp.MustCompile(`for key '(.+?)'`)
			matches := re.FindStringSubmatch(err.Error())
			if len(matches) == 2 {
				campo := matches[1]

				// Ajustamos el nombre del campo para que se vea bonito
				switch campo {
				case "usuarios.correo":
					campo = "correo"
				case "usuarios.nombre":
					campo = "usuario"
				}

				return fmt.Errorf("El %s ya está registrado", campo)
			}
			return fmt.Errorf("Dato duplicado en la base de datos")
		}
	}

	return fmt.Errorf("Error interno: %v", err)
}
