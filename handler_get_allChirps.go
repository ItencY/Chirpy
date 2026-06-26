package main

import "net/http"

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	chirps, err := cfg.db.GetAllChirps(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to get chirps")
		return
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
