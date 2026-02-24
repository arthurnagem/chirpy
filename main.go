package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	assetsHandler := http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("./assets")),
	)

	mux.Handle("/assets/", assetsHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}