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
	router.HandleFunc("/signup", registrationHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", loginHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/profile/{username}", profileHandler).Methods("GET", "OPTIONS")
	http.Handle("/", router)

	stylesServer := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", neuter(stylesServer)))

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


//Соединем фронт и бек логина, делаем редирект после регистрации, 