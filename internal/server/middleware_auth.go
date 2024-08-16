package server

import (
	"log"
	"net/http"

	"github.com/benskia/Blogator/internal/auth"
	"github.com/benskia/Blogator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			log.Println("Error getting ApiKey: ", err)
			respondWithError(w, http.StatusUnauthorized, "Couldn't find API key.")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			log.Println("Error getting user: ", err)
			respondWithError(w, http.StatusUnauthorized, "Invalid ApiKey.")
			return
		}

		handler(w, r, user)
	}
}
