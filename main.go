package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	apiCfg := &apiConfig{}

	// health endpoint
	mux.HandleFunc("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// serve static files
	assetsHandler := http.StripPrefix(
		"/app/",
		http.FileServer(http.Dir(".")),
	)
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(assetsHandler))

	// admin endpoints
	mux.HandleFunc("/admin/metrics", apiCfg.metricsHandler)
	mux.HandleFunc("/admin/reset", apiCfg.resetHandler)

	// API endpoints
	mux.HandleFunc("/api/validate_chirp", apiCfg.validateChirpHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}