package models

import "errors"

var (
	ErrRequiredFirstName = errors.New("nome requerido")
	ErrRequiredLastName  = errors.New("sobrenome requerido")
	ErrRequiredEmail     = errors.New("email requerido")
	ErrRequiredPassword  = errors.New("senha requerido")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	} else {
		return false
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
	if IsEmpty(user.Password) {
		return User{}, ErrRequiredPassword
	}
	return user, nil
}
