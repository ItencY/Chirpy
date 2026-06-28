package main

import (
	"net/http"

	"github.com/itency/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to get token")
		return
	}
	ctx := r.Context()
	err = cfg.db.RevokeRefreshToken(ctx, token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to revoke token")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
