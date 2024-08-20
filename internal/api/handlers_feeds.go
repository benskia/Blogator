package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/benskia/Blogator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedsCreate(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type response struct {
		Feed       `json:"feed"`
		FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Println("Error decoding parameters: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters.")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    u.ID,
	})
	if err != nil {
		log.Println("Error creating feed: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed.")
		return
	}

	feedFollow, err := cfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID:        uuid.New(),
		UserID:    u.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Println("Error following feed: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't follow feed.")
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Feed:       databaseFeedToFeed(feed),
		FeedFollow: databaseFeedFollowToFeedFollow(feedFollow),
	})
}

func (cfg *apiConfig) handleFeedsGetAll(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		log.Println("Error getting feeds: ", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds.")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
