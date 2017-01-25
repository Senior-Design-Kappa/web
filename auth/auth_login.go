package auth

import (
	"bytes"
	"errors"
	"log"

	"golang.org/x/crypto/scrypt"
)

var (
	ErrWrongPassword = errors.New("Wrong password!")
)

func (a Auth) LoginUser(username string, password string) (string, error) {
	row := a.db.QueryRow("SELECT password, salt FROM users WHERE username=?", username)
	var hash []byte
	var salt []byte
	if err := row.Scan(&hash, &salt); err != nil {
		log.Printf("%s\n", err)
		return "", err
	}
	hash2, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 128)
	if err != nil {
		return "", err
	}
	if !bytes.Equal(hash, hash2) {
		return "", ErrWrongPassword
	}
	token, err := a.CreateToken(username)
	return token, err
}
