package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	r := NewRoom()

	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting server on address:", *addr)
	log.Fatal("ListenAndServe:", http.ListenAndServe(*addr, nil))
}
