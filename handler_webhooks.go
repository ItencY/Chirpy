package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/itency/Chirpy/internal/auth"
)

type WebhookRequest struct {
	Event string `json:"event"`
	Data  data   `json:"data"`
}

type data struct {
	UserId uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerWebhooks(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "failed to get api key")
		return
	}
	if cfg.polkaKey != apiKey {
		respondWithError(w, http.StatusUnauthorized, "incorect api key")
		return
	}
	var req WebhookRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	if req.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	ctx := r.Context()
	_, err = cfg.db.UpgradeToChirpyRed(ctx, req.Data.UserId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "failed to update")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
