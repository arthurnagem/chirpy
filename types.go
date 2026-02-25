package main

import "sync/atomic"

type apiConfig struct {
	fileserverHits atomic.Int32
}

type parameters struct {
	Body string `json:"body"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type cleanedChirpResponse struct {
	CleanedBody string `json:"cleaned_body"`
}