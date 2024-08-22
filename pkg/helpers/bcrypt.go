package helpers

import (
	"15min-city/pkg/errs"

	"golang.org/x/crypto/bcrypt"
)

const saltStr = 10

func HashPassword(password string) (string, errs.ErrMessage) {
	bytePass := []byte(password)

	hashed, err := bcrypt.GenerateFromPassword(bytePass, saltStr)

	if err != nil {
		if err == bcrypt.ErrHashTooShort {
			return "", errs.NewBadRequestError("Hash too short")
		}

		return "", errs.NewInternalServerError("Something went wrong")
	}

	return string(hashed), nil
}

func ComparePassword(hashed string, password string) errs.ErrMessage {
	byteHash := []byte(hashed)
	passwordHash := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, passwordHash)

	if err != nil {
		return errs.NewBadRequestError("Wrong email/password")
	}

	return nil
}
