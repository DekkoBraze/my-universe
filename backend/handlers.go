package main

import (
	"html/template"
	"log"
	"net/http"
	"my-universe/database"
	"my-universe/authorization"
	"encoding/base64"
	"github.com/gorilla/mux"
	"encoding/json"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Access-Control-Allow-Origin", "*")		
	type RegistrationData struct{
		Email        			string	`json:"email"`
		Username				string  `json:"username"`
		Password				string	`json:"password"`
		PasswordVerification	string	`json:"passwordVerification"`
	}

	var data RegistrationData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Print("registrationHandler: ", err)
	}
	email := data.Email
	username := data.Username
	password := data.Password
	passwordVerification := data.PasswordVerification

	err = authorization.Registration(email, username, password, passwordVerification)
	if err != nil {
		log.Print("registrationHandler: ", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")

	type LoginData struct{
		Email        			string	`json:"email"`
		Password				string	`json:"password"`
	}

	var data LoginData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Print("registrationHandler: ", err)
	}

	email := data.Email
	password := data.Password
	err = authorization.Login(email, password)
	if err != nil {
		log.Print("loginHandler: ", err)
	}

	encodedValue := base64.StdEncoding.EncodeToString([]byte(email + " : my-universe"))
	database.SetUserSessionValue(email, encodedValue)
	
	cookie := &http.Cookie{
		Name: "my-universe",
		Value: encodedValue, 
		Path: "/", 
		MaxAge: 86400,
	}

	type AnswerLogin struct{
		Username        string	`json:"username"`
	}

	http.SetCookie(w, cookie)
	userBson, err := database.GetProfileBySessionValue(encodedValue)
	if err != nil {
		log.Print("loginHandler: ", err)
	}
	username := userBson.Lookup("username").StringValue()
	answer := AnswerLogin{Username: username}
	json.NewEncoder(w).Encode(answer)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	vars := mux.Vars(r)
	username := vars["username"]
	profileBson, err := database.GetProfileByUsername(username)
	if err != nil {
		log.Print("profileHandler: ", err)
	}
	
	type ProfileData struct{
		Username        string	`json:"username"`
		Age				string  `json:"age"`
		Country			string	`json:"country"`
		Status			string	`json:"status"`
		Description		string	`json:"description"`
	}

	profileData := ProfileData{ Username: profileBson.Lookup("username").StringValue(), 
								  Age: profileBson.Lookup("age").StringValue(), 
								  Country: profileBson.Lookup("country").StringValue(), 
								  Status: profileBson.Lookup("status").StringValue(),
								  Description: profileBson.Lookup("description").StringValue() }

	json.NewEncoder(w).Encode(profileData)
}