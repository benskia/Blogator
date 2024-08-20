package api

import "net/http"

func handleReadiness(w http.ResponseWriter, _ *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, response{Status: "ok"})
}

func handleError(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
