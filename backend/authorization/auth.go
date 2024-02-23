package authorization

import (
	"errors"
	//"github.com/gorilla/sessions"
	"my-universe/database"
	"golang.org/x/crypto/bcrypt"
	"encoding/base64"
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

func Login(email string, password string) (string, error){
	hashedPassword, err := database.GetPasswordByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword[5:], []byte(password))
	if err != nil {
		return "", errors.New("password dosent match")
	}

	encodedValue := base64.StdEncoding.EncodeToString([]byte(email + " : my-universe"))
	database.SetUserSessionValue(email, encodedValue)

	return encodedValue, nil
}