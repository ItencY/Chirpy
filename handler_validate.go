package main

import (
	"encoding/json"
	"net/http"
)

type ChirpRequest struct {
	Body string `json:"body"`
}

type ChirpResponse struct {
	Valid bool `json:"valid"`
}

func handleValidateChirp(w http.ResponseWriter, r *http.Request) {
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
	respondWithJSON(w, http.StatusOK, ChirpResponse{Valid: true})
}
