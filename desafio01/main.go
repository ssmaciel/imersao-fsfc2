package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func filmeViewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("filme.html"))
	
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", filmeViewHandler)
	fmt.Println("Servidor iniciado...")
	http.ListenAndServe(":8000", nil)
}