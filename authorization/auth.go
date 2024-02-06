package authorization

import (
	"errors"
	//"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"my-universe/database"
)

func Registration(email string, username string, password string, passwordVerification string) error{
	if password == passwordVerification {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = database.InsertNewUser(email, username, hashedPassword)
	if err != nil {
		return err
	}
	
	return nil
	} else {
		return errors.New("Password dosen't match") 
	}
}

func Login(email string, password string) {
	
}