package main

import (
	"context"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.env != "dev" {
		w.WriteHeader(http.StatusForbidden)
	} else {
		err := cfg.db.DeleteUsers(context.Background())
		if err != nil {
			log.Fatalf("There was an error deleting all users: %s", err)
		}
		cfg.fileserverHits.Store(0)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hits reset to 0"))
	}
}
