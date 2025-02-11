package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	// replace bad words
	cleanedBody := replaceBadWords(params.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleanedBody,
	})
}

func replaceBadWords(body string) string {
	words := strings.Split(body, " ")
	badWords := []string{"kerfuffle", "sharbert", "fornax"}
	var cleanWords []string
	for _, word := range words {
		isClean := true
		for _, badWord := range badWords {
			if strings.ToLower(word) == badWord {
				isClean = false
				break
			}
		}
		if isClean {
			cleanWords = append(cleanWords, word)
		} else {
			cleanWords = append(cleanWords, "****")
		}
	}
	return strings.Join(cleanWords, " ")
}
