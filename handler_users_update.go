package main

import (
	"encoding/json"
	"net/http"

	"github.com/itency/Chirpy/internal/auth"
	"github.com/itency/Chirpy/internal/database"
)

type UpdateRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (cfg *apiConfig) handlerUsersUpate(w http.ResponseWriter, r *http.Request) {
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
	var req UpdateRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to decode")
		return
	}
	hashPass, err := auth.HashPassword(req.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "failed to hash password")
		return
	}
	ctx := r.Context()
	dbUser, err := cfg.db.UpdateUser(ctx, database.UpdateUserParams{
		Email:          req.Email,
		HashedPassword: hashPass,
		ID:             userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update user")
		return
	}
	user := User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email:     dbUser.Email,
	}
	respondWithJSON(w, http.StatusOK, user)
}
