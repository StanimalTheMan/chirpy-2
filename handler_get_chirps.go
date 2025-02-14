package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	type ChirpResponse struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
		UserID    uuid.UUID `json:"user_id"`
	}

	chirps, err := cfg.db.ListChirps(context.Background())
	chirpResponseFormat := make([]ChirpResponse, len(chirps))
	if err != nil {
		log.Fatalf("There was an error fetching chirps: %s", chirps)
	}

	for i, row := range chirps {
		chirpResponseFormat[i] = ChirpResponse{
			ID:        row.ID,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
			Body:      row.Body,
			UserID:    row.UserID,
		}
	}

	respondWithJSON(w, http.StatusOK, chirpResponseFormat)
}

func (cfg *apiConfig) handlerGetChirp(w http.ResponseWriter, r *http.Request) {
	type ChirpResponse struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Body      string    `json:"body"`
		UserID    uuid.UUID `json:"user_id"`
	}

	chirpID := r.PathValue("chirpID")

	id, err := uuid.Parse(chirpID)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}
	chirp, err := cfg.db.GetChirp(context.Background(), id)
	if err != nil {
		log.Fatalf("There was an error fetching chirps: %s", err)
	}

	respondWithJSON(w, http.StatusOK, ChirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
