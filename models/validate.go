package models

import (
	"errors"

	"github.com/badoux/checkmail"
)

var (
	ErrRequiredFirstName = errors.New("nome requerido")
	ErrRequiredLastName  = errors.New("sobrenome requerido")
	ErrRequiredEmail     = errors.New("email requerido")
	ErrInvalidEmail      = errors.New("email inválido")
	ErrRequiredPassword  = errors.New("senha requerido")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	} else {
		return false
	}
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	} else {
		return true
	}
}

func ValidadeNewUser(user User) (User, error) {
	if IsEmpty(user.FirstName) {
		return User{}, ErrRequiredFirstName
	}
	if IsEmpty(user.LastName) {
		return User{}, ErrRequiredLastName
	}
	if IsEmpty(user.Email) {
		return User{}, ErrRequiredEmail
	}
	if !IsEmail(user.Email) {
		return User{}, ErrInvalidEmail
	}
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}
	return user, nil
}
