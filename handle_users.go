package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benskia/Blogator/internal/database"
	"github.com/google/uuid"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := request{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Error decoding request: ", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request body.")
		return
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	}

	respondWithJSON(w, http.StatusOK, newUser)
}
