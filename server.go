package main

import (
	"log"
	"net/http"
	"my-universe/database"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatal("Init: ", err)
	}

    http.HandleFunc("/", indexHandler)
	http.HandleFunc("/registration", registrationHandler)
	http.HandleFunc("/login", loginHandler)

	stylesServer := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", neuter(stylesServer)))

	err = http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}