package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	const port = "8080"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	err := server.ListenAndServe() // blocking call
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
	log.Printf("serving on port: %v", port)

}
