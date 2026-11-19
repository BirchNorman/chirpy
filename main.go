package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {

	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)

	fileserver := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := fileserver.ListenAndServe() // blocking call
	if err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
	log.Printf("serving on port: %v", port)

}
