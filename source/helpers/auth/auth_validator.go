package auth

type AuthDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a *AuthDTO) ValidUsername() error {
	// length := len(a.Username)
	// if length < 5 || length > 50 {
	// 	return errors.New("el nombre de usuario debe tener entre 5 y 50 caracteres")
	// }
	return nil
}

func (a *AuthDTO) ValidPassword() error {
	// if len(a.Password) < 8 {
	// 	return errors.New("la contraseña debe tener al menos 8 caracteres")
	// }
	// if !strings.ContainsAny(a.Password, "!@#$%^&*()_+-=[]{}|;:,.<>?") {
	// 	return errors.New("la contraseña debe contener al menos un carácter especial")
	// }
	// if !strings.ContainsAny(a.Password, "0123456789") {
	// 	return errors.New("la contraseña debe contener al menos un número")
	// }
	// if !strings.ContainsAny(a.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
	// 	return errors.New("la contraseña debe contener al menos una letra mayúscula")
	// }
	return nil
}
