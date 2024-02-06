package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"my-universe/database"
	"my-universe/authorization"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
        if err != nil {
            log.Fatal("registrationHandler: ", err)
        }

		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")
		passwordVerification := r.FormValue("password_verification")

		err = authorization.Registration(email, username, password, passwordVerification)
		if err != nil {
            log.Fatal("registrationHandler: ", err)
        }
		
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		tmpl, _ := template.ParseFiles("templates/registration.html")
		tmpl.Execute(w, nil)
	}
}

//Доделать сокрытие структуры и файлов
func neuter(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
            http.NotFound(w, r)
            return
        }

        next.ServeHTTP(w, r)
    })
}
 
func main() {
	err := database.Init()
	if err != nil {
		log.Fatal("Init: ", err)
	}

    http.HandleFunc("/", indexHandler)
	http.HandleFunc("/registration", registrationHandler)

	stylesServer := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", neuter(stylesServer)))

	err = http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}