package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)
 
func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, nil)
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
    http.HandleFunc("/", indexHandler)

	stylesServer := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", neuter(stylesServer)))

	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}