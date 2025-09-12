package user_model

type UserSanitizer struct {
	ID              int
	PrimerNombre    string
	SegundoNombre   *string
	PrimerApellido  string
	SegundoApellido *string
	Matricula       string
	Correo          string
	Rol             *string
}
