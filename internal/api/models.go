package api

import (
	"time"

	"github.com/benskia/Blogator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.Apikey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(f database.Feed) Feed {
	return Feed{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		URL:       f.Url,
		UserID:    f.UserID,
	}
}

func databaseFeedsToFeeds(f []database.Feed) []Feed {
	feeds := make([]Feed, len(f))
	for i, feed := range f {
		feeds[i] = databaseFeedToFeed(feed)
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(f database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		UserID:    f.UserID,
		FeedID:    f.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(f []database.FeedFollow) []FeedFollow {
	feeds := make([]FeedFollow, len(f))
	for i, feed := range f {
		feeds[i] = databaseFeedFollowToFeedFollow(feed)
	}
	return feeds
}
