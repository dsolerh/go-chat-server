package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
        <html>
           <head>
             <title>Chat</title>
           </head>
           <body>
             Let's chat!
           </body>
        </html>
        `))
	})

	// start the web server
	log.Fatal("ListenAndServe:", http.ListenAndServe(":8080", nil))
}
