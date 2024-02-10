package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	log.Printf("listening on port %s", port)

	c := NewHandler()
	http.HandleFunc("/test", c.Handle)

	log.Println("starting HTTP server")
	if err := http.ListenAndServe(":"+port, nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server closed: %v", err)
	}
}
