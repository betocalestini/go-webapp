package auth

import (
	"errors"
	"go-webapp/models"
	"go-webapp/utils"
)

var (
	ErrEmailNotFound   = errors.New("email não encontrado")
	ErrInvalidPassword = errors.New("senha inválida")
	ErrEmptyFields     = errors.New("preencha todos os campos")
)

func Signin(email, password string) (models.User, error) {
	err := validateFields(email, password)
	if err != nil {
		return models.User{}, err
	}
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, ErrEmailNotFound
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}

func validateFields(email, password string) error {
	if models.IsEmpty(email) || models.IsEmpty(password) {
		return ErrEmptyFields
	}
	if !models.IsEmail(email) {
		return models.ErrInvalidEmail
	}
	return nil
}
