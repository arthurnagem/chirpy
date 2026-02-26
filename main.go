package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "github.com/arthurnagem/chirpy/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	
	dbQueries := database.New(db)

	
	apiCfg := &apiConfig{
		Queries: dbQueries,
		Platform: os.Getenv("PLATFORM"),
	}

	mux := http.NewServeMux()

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
	mux.HandleFunc("/api/chirps", apiCfg.createChirpHandler)

	mux.HandleFunc("/api/users", apiCfg.createUserHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}