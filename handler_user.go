package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type UserRequest struct {
	Email string `json:"email"`
}

func (cfg *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	ctx := r.Context()
	dbUser, err := cfg.db.CreateUser(ctx, req.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed created user")
		return
	}
	user := User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}
	respondWithJSON(w, http.StatusCreated, user)
}
