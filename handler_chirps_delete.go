package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/itency/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerDeleteChirps(w http.ResponseWriter, r *http.Request) {
	acessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "failed to get token")
		return
	}
	userID, err := auth.ValidateJWT(acessToken, cfg.jwt)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
		return
	}
	chirpID := r.PathValue("chirpID")
	chirpIDParse, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to parse string to uuid")
		return
	}
	ctx := r.Context()
	chirp, err := cfg.db.GetChirpByID(ctx, chirpIDParse)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "failed to get chirp by ID")
		return
	}
	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Incorrect access rights")
		return
	}
	_, err = cfg.db.DeleteChirpByID(ctx, chirpIDParse)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "chipr not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
