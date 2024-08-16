package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benskia/Blogator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Error decoding parameters: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters.")
		return
	}

	if len(params.Name) == 0 {
		respondWithError(w, http.StatusBadRequest, "Name parameter was empty.")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Println("Error creating user: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user.")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (cfg *apiConfig) handleUsersGetOne(w http.ResponseWriter, r *http.Request, u database.User) {
	respondWithJSON(w, http.StatusOK, u)
}
