package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	// setup gomniauth
	gomniauth.SetSecurityKey("SECRET KEY")
	gomniauth.WithProviders(
		facebook.New("key", "secret", "http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret", "http://localhost:8080/auth/callback/github"),
		google.New("825102912466-bcqv49ighbhpdkeqgmb5v1a6buhcgjlv.apps.googleusercontent.com", "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", "http://localhost:8080/auth/callback/google"),
	)

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
