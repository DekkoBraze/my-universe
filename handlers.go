package main

import (
	"html/template"
	"log"
	"net/http"
	"my-universe/database"
	"my-universe/authorization"
	"encoding/base64"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
        if err != nil {
            log.Print("registrationHandler: ", err)
        }

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")
		passwordVerification := r.FormValue("password_verification")

		err = authorization.Registration(email, username, password, passwordVerification)
		if err != nil {
            log.Print("registrationHandler: ", err)
        }
		
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	} else {
		tmpl, _ := template.ParseFiles("templates/registration.html")
		tmpl.Execute(w, nil)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
        if err != nil {
            log.Print("loginHandler: ", err)
        }

		email := r.FormValue("email")
		password := r.FormValue("password")

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

		http.SetCookie(w, cookie)
		
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.Execute(w, nil)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	profileBson, err := database.GetProfile(username)
	if err != nil {
		log.Print("profileHandler: ", err)
	}
	
	
	type TemplateData struct{
		Username        string
		Age				string
		Country			string
		Status			string
		Description		string
	}

	tmpl, _ := template.ParseFiles("templates/profile.html")
	templateData := TemplateData{ Username: profileBson.Lookup("username").StringValue(), 
								  Age: profileBson.Lookup("age").StringValue(), 
								  Country: profileBson.Lookup("country").StringValue(), 
								  Status: profileBson.Lookup("status").StringValue(),
								  Description: profileBson.Lookup("description").StringValue() }
	tmpl.Execute(w, templateData)
}