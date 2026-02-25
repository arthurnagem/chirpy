package main

import (
	"sync/atomic"
	
	"github.com/arthurnagem/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	Queries        *database.Queries
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

type createUserRequest struct {
	Email string `json:"email"`
}