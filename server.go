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

    router.HandleFunc("/", indexHandler)
	router.HandleFunc("/registration", registrationHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/profile/{username}", profileHandler)
	http.Handle("/", router)

	stylesServer := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", neuter(stylesServer)))

	err = http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}