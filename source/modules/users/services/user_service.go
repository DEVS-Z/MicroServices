package user_service

import (
	roles_model "main/source/modules/roles/model"
	roles_services "main/source/modules/roles/services"
	user_model "main/source/modules/users/models"

	base_service "github.com/miqueaz/FrameGo/pkg/base/service"
)

type UserService struct {
	base_service.Service[user_model.UserStruct]
}

func (s *UserService) ReadWithOutSanitizer(filter map[string]any) ([]user_model.UserStruct, error) {
	return s.Service.Read(filter)
}

func (s *UserService) Read(filters map[string]any) ([]user_model.UserSanitizer, error) {
	datas, err := s.Service.Read(filters)
	if err != nil {
		return nil, err
	}

	var sanitizers []user_model.UserSanitizer

	for _, data := range datas {
		rol, err := roles_services.Service.ReadOne(data.Rol)
		if err != nil {
			rol = roles_model.RolesStruct{}
			continue
		}

		sanitizer := user_model.UserSanitizer{
			ID:              data.ID,
			PrimerNombre:    data.PrimerNombre,
			SegundoNombre:   data.SegundoNombre,
			PrimerApellido:  data.PrimerApellido,
			SegundoApellido: data.SegundoApellido,
			Matricula:       data.Matricula,
			Correo:          data.Correo,
			Rol:             &rol.Nombre,
		}
		sanitizers = append(sanitizers, sanitizer)
	}

	return sanitizers, nil
}

func (s *UserService) ReadOne(id int) (*user_model.UserSanitizer, error) {
	// Implementation for reading user data

	datas, err := s.Service.Read(map[string]any{"ID": id})

	if len(datas) <= 0 {
		return nil, err
	}

	data := datas[0]

	role, err := roles_services.Service.ReadOne(data.Rol)

	if err != nil {
		role = roles_model.RolesStruct{}
	}

	sanitizer := user_model.UserSanitizer{
		ID:              data.ID,
		PrimerNombre:    data.PrimerNombre,
		SegundoNombre:   data.SegundoNombre,
		PrimerApellido:  data.PrimerApellido,
		SegundoApellido: data.SegundoApellido,
		Matricula:       data.Matricula,
		Correo:          data.Correo,
		Rol:             &role.Nombre,
	}

	return &sanitizer, nil
}

func (s *UserService) Insert(user user_model.UserStruct) (*user_model.UserSanitizer, error) {

	user.Contrasena = user.Matricula + "@12345"

	newUser, err := s.Service.Insert(user)
	if err != nil {
		return nil, err
	}

	usr, err := s.ReadOne(newUser.ID)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

var Service = base_service.NewService[UserService](*user_model.Model)
