package authorization

import (
	"errors"
	//"github.com/gorilla/sessions"
	"my-universe/database"
	"golang.org/x/crypto/bcrypt"
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
		return errors.New("password dosen't match") 
	}
}

func Login(email string, password string) error{
	hashedPassword, err := database.GetPasswordByEmail(email)
	//hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword[5:], []byte(password))
	if err != nil {
		return errors.New("password dosent match")
	}

	return nil
}