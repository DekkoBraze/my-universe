package main

import (
	"log"
	"net/http"
	"my-universe/database"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatal("Init: ", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprint(w, "API Index Page")
    })
	subrouter := router.PathPrefix("/api/").Subrouter()
	subrouter.HandleFunc("/signup", registrationHandler).Methods("POST", "OPTIONS")
	subrouter.HandleFunc("/login", loginHandler).Methods("POST", "OPTIONS")
	subrouter.HandleFunc("/getUser", getCurrentUser).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/profile/{username}", profileHandler).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/createItem", addItem).Methods("POST", "OPTIONS")
	subrouter.HandleFunc("/getUserItems/{username}", getUserItems).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/deleteUserItem/{username}/{itemId}", deleteUserItem).Methods("DELETE", "OPTIONS")
	subrouter.HandleFunc("/getRawgKey", rawgHandler).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/getTmdbKey", tmdbHandler).Methods("GET", "OPTIONS")
	subrouter.HandleFunc("/getLastfmKey", lastfmHandler).Methods("GET", "OPTIONS")
	http.Handle("/", router)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

