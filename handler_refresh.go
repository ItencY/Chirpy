package main

import (
	"net/http"
	"time"

	"github.com/itency/Chirpy/internal/auth"
)

type RefreshResponse struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to get token")
		return
	}
	ctx := r.Context()
	userRefreshToken, err := cfg.db.GetUserFromRefreshToken(ctx, token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid refresh token")
		return
	}
	duration := time.Hour
	jwtToken, err := auth.MakeJWT(userRefreshToken.ID, cfg.jwt, duration)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to create jwt")
		return
	}
	respondWithJSON(w, http.StatusOK, RefreshResponse{Token: jwtToken})
}
