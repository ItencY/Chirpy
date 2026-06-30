package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/itency/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	authorIDStr := r.URL.Query().Get("author_id")
	ctx := r.Context()
	var chirps []database.Chirp
	var err error
	if authorIDStr == "" {
		chirps, err = cfg.db.GetAllChirps(ctx)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "failed to get chirps")
			return
		}
	} else {
		authorID, err := uuid.Parse(authorIDStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "failed to parse")
			return
		}
		chirps, err = cfg.db.GetChirpByAuthorID(ctx, authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "failed to get chirp")
			return
		}
	}
	var allChirps []Chrips
	for _, chirp := range chirps {
		allChirps = append(allChirps, Chrips{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, allChirps)
}
