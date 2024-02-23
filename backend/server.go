package main

import (
	"log"
	"net/http"
	"my-universe/database"
	"github.com/gorilla/mux"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatal("Init: ", err)
	}

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/").Subrouter()
	subrouter.HandleFunc("/signup", registrationHandler).Methods("POST", "OPTIONS")
	subrouter.HandleFunc("/login", loginHandler).Methods("POST", "OPTIONS")
	subrouter.HandleFunc("/getUser", getCurrentUser).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/profile/{username}", profileHandler).Methods("GET", "OPTIONS")
	http.Handle("/", router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

