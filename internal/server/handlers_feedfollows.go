package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benskia/Blogator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedFollowsAdd(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Error decoding parameters: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters.")
		return
	}

	feedID, err := uuid.Parse(params.FeedID)
	if err != nil {
		log.Println("Error parsing feed ID: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse feed ID.")
		return
	}

	feedFollow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		UserID:    u.ID,
		FeedID:    feedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Println("Error following feed: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow feed.")
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollow)
}

func (cfg *apiConfig) handleFeedFollowsDelete(w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollowIDPathValue := r.PathValue("feedFollowID")
	if feedFollowIDPathValue == "" {
		respondWithError(w, http.StatusBadRequest, "No feedFollowID provided.")
		return
	}

	feedFollowID, err := uuid.Parse(feedFollowIDPathValue)
	if err != nil {
		log.Println("Error parsing string to UUID: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse feedFollowID.")
		return
	}

	feedFollow, err := cfg.DB.GetFeedFollowByID(r.Context(), feedFollowID)
	if err != nil {
		log.Println("Error getting feed follow: ", err)
		respondWithError(w, http.StatusBadRequest, "Couldn't find feedFollow.")
		return
	}

	if feedFollow.UserID != u.ID {
		respondWithError(w, http.StatusUnauthorized, "User doesn't own this feedFollow.")
		return
	}

	err = cfg.DB.UnfollowFeed(r.Context(), feedFollowID)
	if err != nil {
		log.Println("Error unfollowing feed: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't unfollow feed.")
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) handleFeedFollowsGetAllByUser(w http.ResponseWriter, r *http.Request, u database.User) {
	feedFollows, err := cfg.DB.GetAllFeedFollowsByUserID(r.Context(), u.ID)
	if err != nil {
		log.Println("Error getting feedFollows: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feedFollows.")
		return
	}

	respondWithJSON(w, http.StatusOK, feedFollows)
}
