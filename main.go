package main

import (
	"html/template"
	"net/http"
	"fmt"
)

type Page struct{
	value string
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/index.html"))
	p:=Page{value: "test"}
	err := templates.ExecuteTemplate(w, "index.html", p)
	fmt.Println("Server started  to port 8080")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
