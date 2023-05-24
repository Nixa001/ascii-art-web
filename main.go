package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.IndexHandle)
	http.HandleFunc("/ascii-art", handlers.HomeHandler)
	fmt.Println("Server started on port", port)
	fmt.Println("Open link -> (http://localhost:8080/)")
	http.ListenAndServe(port, nil)
}
