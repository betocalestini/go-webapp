package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/badoux/checkmail"
)

var (
	ErrRequiredFirstName = errors.New("nome requerido")
	ErrRequiredLastName  = errors.New("sobrenome requerido")
	ErrRequiredEmail     = errors.New("email requerido")
	ErrInvalidEmail      = errors.New("email inválido")
	ErrRequiredPassword  = errors.New("senha requerido")
	ErrMaxLimit          = errors.New("ultrapassou o limite máximo de caracteres")
	ErrDuplicateKeyEmail = errors.New("o email informado já possui cadastro")
)

func IsEmpty(attr string) bool {
	if attr == "" {
		return true
	} else {
		return false
	}
}

func Trim(attr string) string {
	return strings.TrimSpace(attr)
}

func IsEmail(email string) bool {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return false
	} else {
		return true
	}
}

func Max(attr string, lim int) bool {
	if len(attr) <= lim {
		return true
	} else {
		return false
	}
}

func ValidateLimitFields(user User) (User, error) {
	if !Max(user.FirstName, 15) || !Max(user.LastName, 20) || !Max(user.Email, 40) || !Max(user.Password, 100) {
		return user, ErrMaxLimit
	}
	return user, nil
}

func UniqueEmail(email string) (bool, error) {
	con := Connect()
	defer con.Close()
	sql := "select count(email) from users where email = $1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return false, err
	}
	defer rs.Close()
	var count int64
	if rs.Next() {
		err := rs.Scan(&count)
		if err != nil {
			return false, err
		}
		if count > 0 {
			return false, ErrDuplicateKeyEmail
		}
	}

	return true, nil
}

func ValidadeNewUser(user User) (User, error) {
	_, err := UniqueEmail(user.Email)
	if err != nil {
		return User{}, err
	}
	user, err = ValidateLimitFields(user)
	if err != nil {
		return user, err
	}
	user.FirstName = Trim(user.FirstName)
	user.LastName = Trim(user.LastName)
	user.Email = Trim(strings.ToLower(user.Email))

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

func Count(table string) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := fmt.Sprintf("select count(*) from %s", table)
	var count int64
	err := con.QueryRow(sql).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func LoadData() (int64, int64, error) {
	allProducts, err := Count("products")
	if err != nil {
		return 0, 0, err
	}
	allUsers, err := Count("users")
	if err != nil {
		return 0, 0, err
	}
	return allProducts, allUsers, nil
}
