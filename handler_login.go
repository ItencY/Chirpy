package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/itency/Chirpy/internal/auth"
	"github.com/itency/Chirpy/internal/database"
)

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	User
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	duration := time.Hour
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
	makeRefreshToken := auth.MakeRefreshToken()
	refreshToken, err := cfg.db.CreateRefreshToken(ctx, database.CreateRefreshTokenParams{
		Token:     makeRefreshToken,
		UserID:    userEmail.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 60),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to save refresh token")
		return
	}
	user := User{
		ID:          userEmail.ID,
		CreatedAt:   userEmail.CreatedAt,
		UpdatedAt:   userEmail.UpdatedAt,
		Email:       userEmail.Email,
		IsChirpyRed: userEmail.IsChirpyRed.Bool,
	}
	respondWithJSON(w, http.StatusOK, LoginResponse{User: user, Token: token, RefreshToken: refreshToken.Token})
}
