package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting server on port: 8080")
	log.Fatal("ListenAndServe:", http.ListenAndServe(":8080", nil))
}
