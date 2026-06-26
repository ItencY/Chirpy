package main

import (
	"encoding/json"
	"net/http"

	"github.com/itency/Chirpy/internal/auth"
)

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	ctx := r.Context()
	userEmail, err := cfg.db.GetUserByEmail(ctx, req.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	password, err := auth.CheckPasswordHash(req.Password, userEmail.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	if !password {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}
	user := User{
		ID:        userEmail.ID,
		CreatedAt: userEmail.CreatedAt,
		UpdatedAt: userEmail.UpdatedAt,
		Email:     userEmail.Email,
	}
	respondWithJSON(w, http.StatusOK, user)
}
