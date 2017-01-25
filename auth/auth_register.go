package auth

import (
	"crypto/rand"
	"database/sql"
	"errors"

	"golang.org/x/crypto/scrypt"
)

var (
	ErrUserExists  = errors.New("User already exists!")
	ErrEmailExists = errors.New("Email already in use!")
)

func (a Auth) RegisterNewUser(username string, password string, email string) error {
	// Making sure user/email doesn't collide
	row := a.db.QueryRow("SELECT username FROM users WHERE username=?", username)
	if err := row.Scan(); err != sql.ErrNoRows {
		return ErrUserExists
	}
	row := a.db.QueryRow("SELECT email FROM users WHERE email=?", email)
	if err := row.Scan(); err != sql.ErrNoRows {
		return ErrEmailExists
	}

	// Cryptography
	salt := make([]byte, 128)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 128)
	if err != nil {
		return err
	}

	// Inserting into DB
	stmt, err := a.db.Prepare("INSERT INTO USERS (email, password, username, salt) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(email, hash, username, salt)
	if err != nil {
		return err
	}

	return nil
}
