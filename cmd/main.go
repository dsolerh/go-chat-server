package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})

	// start the web server
	log.Fatal("ListenAndServe:", http.ListenAndServe(":8080", nil))
}
