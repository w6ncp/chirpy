package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/w6ncp/chirpy/internal/database"
)

func (cfg *apiConfig) handlerListChirps(w http.ResponseWriter, r *http.Request) {
	var err error
	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	}

	var chirps []database.Chirp
	if authorID != uuid.Nil {
		chirps, err = cfg.db.GetChirpsByUser(r.Context(), authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not get chirps", err)
			return
		}
	} else {
		chirps, err = cfg.db.GetChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not get chirps", err)
			return
		}
	}

	listChirp := []Chirp{}

	for _, dbChirp := range chirps {
		listChirp = append(listChirp, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		})
	}

	sortString := r.URL.Query().Get("sort")
	if sortString == "desc" {
		sort.Slice(listChirp, func(i, j int) bool {
			return listChirp[j].CreatedAt.Before(listChirp[i].CreatedAt)
		})
	}

	respondWithJSON(w, http.StatusOK, listChirp)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}
