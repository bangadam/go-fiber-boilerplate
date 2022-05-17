package domain

import (
	"crypto/md5"
	"encoding/hex"

	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Username string `json:"username" example:"admin1"`
	Password string `json:"password" example:"admin1"`
}

// create validator for login request
func (l LoginRequest) LoginValidation() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username, validation.Required),
		validation.Field(&l.Password, validation.Required),
	)
}

// hashed md5 password
func (l LoginRequest) MD5Password() string {
	hash := md5.New()
	hash.Write([]byte(l.Password))
	return hex.EncodeToString(hash.Sum(nil))
}