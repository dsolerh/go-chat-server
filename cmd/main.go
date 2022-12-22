package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// start the web server
	log.Println("Starting server on port: 8080")
	log.Fatal("ListenAndServe:", http.ListenAndServe(":8080", nil))
}
