package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTCreationAndValidation(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	returnedID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatal(err)
	}

	if returnedID != userID {
		t.Errorf("expected %v, got %v", userID, returnedID)
	}
}

func TestExpiredJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Error("expected error for expired token")
	}
}

func TestJWTWrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateJWT(token, "wrong-secret")
	if err == nil {
		t.Error("expected error for wrong secret")
	}
}