package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailParams struct {
	To         string `json:"to"`
	Subject    string `json:"subject"`
	Matricula  string `json:"matricula"`
	Contrasena string `json:"contrasena"`
}

var url = "http://mailer:3000"

func SendMail(matricula string, contrasena string, email string) error {
	// Codificar el cuerpo como JSON
	emailParams := EmailParams{
		To:         email,
		Subject:    "Bienvenida a QubeFlex Droply",
		Matricula:  matricula,
		Contrasena: contrasena,
	}
	body, err := json.Marshal(emailParams)
	if err != nil {
		return fmt.Errorf("error al codificar el cuerpo JSON: %v", err)
	}

	// Construir la solicitud POST
	resp, err := http.Post(fmt.Sprintf("%s/send", url), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Verificar código de respuesta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("respuesta no exitosa del servidor: %s", resp.Status)
	}

	return nil
}
