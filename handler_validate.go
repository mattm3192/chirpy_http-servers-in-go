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
		Cleaned_Body string `json:"cleaned_body"`
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

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanBody(params.Body, badWords)

	respondWithJSON(w, http.StatusOK, returnVals{Cleaned_Body: cleaned})
}

func getCleanBody(body string, badWords map[string]struct{}) string {
	contents := strings.Split(body, " ")
	for i, word := range contents {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			contents[i] = "****"
		}
	}
	return strings.Join(contents, " ")
}
