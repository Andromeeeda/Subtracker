package hashpassword

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",errors.New("password generation error")
	}

	return string(hashPassword),nil 
}

func CheckHashPassword(hashPassword string,password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password)); err != nil {
		return errors.New("inccorect password!")
	}

	return nil 
}
