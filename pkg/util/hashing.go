package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword, password string) error {
	// Compare the hashed password with the password
	//return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return errors.New("password is incorrect")
	}

	if err != nil {
		return err
	}

	return nil
}
