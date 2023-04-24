package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	value string
}

func main() {
	http.HandleFunc("/home", homeHandler)

	fmt.Println("Server started  to port 8080")
	fmt.Println("Open link -> http://localhost:8080/home")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseFiles("templates/index.html"))
	p := Page{value: "test"}
	err := templates.ExecuteTemplate(w, "index.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
