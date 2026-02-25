package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func cleanProfanity(body string) string {
	profanities := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(body, " ")

	for i, w := range words {
		lower := strings.ToLower(w)
		for _, p := range profanities {
			if lower == p {
				words[i] = "****"
				break
			}
		}
	}

	return strings.Join(words, " ")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, errorResponse{Error: msg})
}