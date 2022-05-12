package auth

import (
	"errors"
	"go-webapp/models"
	"go-webapp/utils"
)

var (
	ErrInvalidEmail    = errors.New("email inválido")
	ErrInvalidPassword = errors.New("senha inválida")
)

func Signin(email, password string) (models.User, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, ErrInvalidEmail
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{}, ErrInvalidPassword
	}
	return user, nil
}
