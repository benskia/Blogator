package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Println("Error encoding payload: ", err)
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, statusCode, errorResponse{Error: message})
}
