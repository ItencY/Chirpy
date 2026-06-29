package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/itency/Chirpy/internal/auth"
	"github.com/itency/Chirpy/internal/database"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `json:"email"`
	IsChirpyRed bool      `json:"is_chirpy_red"`
}

type UserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (cfg *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to hash password")
		return
	}
	ctx := r.Context()
	dbUser, err := cfg.db.CreateUser(ctx, database.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed created user")
		return
	}
	user := User{
		ID:          dbUser.ID,
		CreatedAt:   dbUser.CreatedAt,
		UpdatedAt:   dbUser.UpdatedAt,
		Email:       dbUser.Email,
		IsChirpyRed: dbUser.IsChirpyRed.Bool,
	}
	respondWithJSON(w, http.StatusCreated, user)
}
