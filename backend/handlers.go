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

type Response struct{
	Message        			string	`json:"message"`
	Username				string	`json:"username"`
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {	
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")		
	type RegistrationData struct{
		Email        			string	`json:"email"`
		Username				string  `json:"username"`
		Password				string	`json:"password"`
		DateOfBirth				string	`json:"dateOfBirth"`
		Gender					string	`json:"gender"`
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

	err = authorization.Registration(data.Email, data.Username, data.Password, 
		data.DateOfBirth, data.Gender, data.Country, data.Status, data.Description)
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
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")
	userFound := false
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

func rawgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")

	const rawgApiKey = "26ed1d032db149269419f1eadbf3d3f7"
	var response = Response{Message: rawgApiKey}
	json.NewEncoder(w).Encode(response)
}

func tmdbHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")

	const tmdbApiKey = "3e1ae42bc540ef6a25740c1638eed8fd"
	var response = Response{Message: tmdbApiKey}
	json.NewEncoder(w).Encode(response)
}

func lastfmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")

	const lastfmKey = "1babbce302adbb3994ecdafe3366694a"
	var response = Response{Message: lastfmKey}
	json.NewEncoder(w).Encode(response)
}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	w.Header().Set("Content-Type", "application/json")

	type ItemData struct{
		Username        			string	`json:"username"`
		ItemType					string 	`json:"itemType"`
		ItemId						string  `json:"itemId"`
		ItemName					string 	`json:"itemName"`
		ItemImage					string  `json:"itemImage"`
		Rating						int		`json:"rating"`
		Comment						string	`json:"comment"`
	}

	var data ItemData

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Print("addItem: ", err)
	}

	err = database.AddItem(data.Username, data.ItemType, data.ItemId, data.ItemName, data.ItemImage, data.Rating, data.Comment)
	if err != nil {
		log.Print("addItem: ", err)
		var response = Response{Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	} else {
		var response = Response{Message: "OK"}
		json.NewEncoder(w).Encode(response)
	}
}

func getUserItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)

	vars := mux.Vars(r)
	username := vars["username"]
	items, err := database.GetUserItems(username)
	if err != nil {
		log.Print("getUserItems: ", err)
	}

	json.NewEncoder(w).Encode(items)
}

func deleteUserItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", reactUrl)
	
	vars := mux.Vars(r)
	username := vars["username"]
	itemId := vars["itemId"]

	err := database.DeleteUserItem(username, itemId) //не находит айтем
	if err != nil {
		log.Print("deleteUserItem: ", err)
		var response = Response{Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	} else {
		var response = Response{Message: "OK"}
		json.NewEncoder(w).Encode(response)
	}
}