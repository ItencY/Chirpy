package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/itency/Chirpy/internal/auth"
)

type LoginRequest struct {
	Password         string `json:"password"`
	Email            string `json:"email"`
	ExpiresInSeconds int    `json:"expires_in_seconds"`
}

type LoginResponse struct {
	User
	Token string `json:"token"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	duration := determineExpiration(req.ExpiresInSeconds)
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
	token, err := auth.MakeJWT(userEmail.ID, cfg.jwt, duration)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create jwt")
		return
	}
	user := User{
		ID:        userEmail.ID,
		CreatedAt: userEmail.CreatedAt,
		UpdatedAt: userEmail.UpdatedAt,
		Email:     userEmail.Email,
	}
	respondWithJSON(w, http.StatusOK, LoginResponse{User: user, Token: token})
}

func determineExpiration(reqExpires int) time.Duration {
	maxDuration := 1 * time.Hour
	if reqExpires <= 0 {
		return maxDuration
	}
	reqDuration := time.Duration(reqExpires) * time.Second
	if reqDuration > maxDuration {
		return maxDuration
	}
	return reqDuration
}
