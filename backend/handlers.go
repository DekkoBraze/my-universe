package main

import (
	"log"
	"net/http"
	"my-universe/database"
	"my-universe/authorization"
	"github.com/gorilla/mux"
	"encoding/json"
)

var reactUrl = "http://localhost:3000"

func registrationHandler(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")		
	type RegistrationData struct{
		Email        			string	`json:"email"`
		Username				string  `json:"username"`
		Password				string	`json:"password"`
		DateOfBirth				string	`json:"dateOfBirth"`
		Country					string	`json:"country"`
		Status					string	`json:"status"`
		Description				string	`json:"description"`
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
	dateOfBirth := data.DateOfBirth
	country := data.Country
	status := data.Status
	description := data.Description
	
	type Response struct{
		Message        			string	`json:"message"`
	}

	err = authorization.Registration(email, username, password, dateOfBirth, country, status, description)
	if err != nil {
		var response = Response{Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	} else {
		var response = Response{Message: "OK"}
		json.NewEncoder(w).Encode(response)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")

	type LoginData struct{
		Email        			string	`json:"email"`
		Password				string	`json:"password"`
	}

	var data LoginData
	
	type Response struct{
		Message        			string	`json:"message"`
		Username				string	`json:"username"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Print("registrationHandler: ", err)
	}

	email := data.Email
	password := data.Password
	encodedValue, err := authorization.Login(email, password)
	if err != nil {
		var response = Response{Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	
	if encodedValue != "" {
			cookie := &http.Cookie{
			Name: "my-universe",
			Value: encodedValue, 
			Path: "/", 
			MaxAge: 86400,
		}
		http.SetCookie(w, cookie)
	
		userBson, err := database.GetProfileBySessionValue(encodedValue)
		if err != nil {
			log.Print("loginHandler: ", err)
		}
		
		username := userBson.Lookup("username").StringValue()
		answer := Response{Username: username, Message: "OK"}
		json.NewEncoder(w).Encode(answer)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	vars := mux.Vars(r)
	username := vars["username"]
	profileBson, err := database.GetProfileByUsername(username)
	if err != nil {
		log.Print("profileHandler: ", err)
	}
	
	profileJson := profileFromBsonToJson(profileBson)
	json.NewEncoder(w).Encode(profileJson)
}

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	userFound := false
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	for _, cookie := range r.Cookies() {
    if cookie.Name == "my-universe" {
		profileBson, err := database.GetProfileBySessionValue(cookie.Value)
		if err != nil {
			log.Print("profileHandler: ", err)
		}
		profileJson := profileFromBsonToJson(profileBson)
		json.NewEncoder(w).Encode(profileJson)
		userFound = true
    }
}
	if (!userFound) {json.NewEncoder(w).Encode("USER_NOT_LOGGED")}
}