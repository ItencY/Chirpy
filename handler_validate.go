package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ChirpRequest struct {
	Body string `json:"body"`
}

type ChirpResponse struct {
	CleanedBody string `json:"cleaned_body"`
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
	respondWithJSON(w, http.StatusOK, ChirpResponse{CleanedBody: cleanBody})
}
