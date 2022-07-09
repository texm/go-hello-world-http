package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	var port string

	if port = os.Getenv("PORT"); port == "" {
		log.Fatalf("Environment variable 'PORT' is not set, exiting...")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := 200
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			code = 404
		} else {
			fmt.Fprintf(w, "Hello, world!\n")
		}
		log.Printf("%d %s %s %s\n", code, r.Proto, r.Method, r.URL.Path)
	})

	addr := net.JoinHostPort("0.0.0.0", port)

	log.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Closed server")
}
