package main

import (
	"sync/atomic"
	"time"
	"github.com/google/uuid"
	
	"github.com/arthurnagem/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	Queries        *database.Queries
	Platform	   string
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
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

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}