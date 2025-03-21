package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/w6ncp/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Could not find API Key", err)
		return
	}
	if apiKey != cfg.polkaKey {
		err = fmt.Errorf("api key mismatch")
		respondWithError(w, http.StatusConflict, "Could not validate API Key", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		err = fmt.Errorf("unable to handle event")
		respondWithError(w, http.StatusNoContent, "Unauthorized event", err)
		return
	}

	_, err = cfg.db.UpgradeByUser(r.Context(), params.Data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Count not find user", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not upgrade user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
