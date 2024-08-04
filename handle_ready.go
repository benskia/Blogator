package main

import "net/http"

func readinessHandle(w http.ResponseWriter, _ *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, response{Status: "ok"})
}

func errHandle(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusOK, "Internal Server Error")
}
