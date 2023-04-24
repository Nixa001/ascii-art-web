package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/home", handlers.HomeHandler)
	fmt.Println("Server started on port", port)
	fmt.Println("Open link -> (http://localhost:8080/home)")
	http.ListenAndServe(port, nil)
}
