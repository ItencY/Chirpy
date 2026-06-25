package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/itency/Chirpy/internal/database"
)

type Chrips struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

type ChirpRequest struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handleCreateChirp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
	var req ChirpRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}
	if len(req.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	words := strings.Split(req.Body, " ")
	cleanedWords := make([]string, len(words))
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if lowerWord == "kerfuffle" || lowerWord == "sharbert" || lowerWord == "fornax" {
			cleanedWords[i] = "****"
		} else {
			cleanedWords[i] = word
		}
	}
	cleanBody := strings.Join(cleanedWords, " ")
	ctx := r.Context()
	chirp, err := cfg.db.CreateChirp(ctx, database.CreateChirpParams{
		Body:   cleanBody,
		UserID: req.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed create chirp")
		return
	}
	chirps := Chrips{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}
	respondWithJSON(w, http.StatusCreated, chirps)
}
