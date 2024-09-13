package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	const logofilepath = "/assets"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/assets", http.FileServer(http.Dir(logofilepath)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s %s on port: %s\n", logofilepath, filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
