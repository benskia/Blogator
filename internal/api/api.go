package api

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/benskia/Blogator/internal/database"
	_ "github.com/lib/pq" // The psql driver is needed for its side-effects.
)

type EnvVars struct {
	DbURL string
	Port  string
}

type apiConfig struct {
	DB             *database.Queries
	fetch_count    int
	fetch_interval time.Duration
}

type parsedXML struct {
	InnerXML []byte `xml:"innerxml"`
}

func StartBlogator(env EnvVars) {
	db, err := sql.Open("postgres", env.DbURL)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB:          dbQueries,
		fetch_count: 10,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handleReadiness)
	mux.HandleFunc("GET /v1/err", handleError)

	mux.HandleFunc("POST /v1/users", apiCfg.handleUsersCreate)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.handleUsersGetOne))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.handleFeedsCreate))
	mux.HandleFunc("GET /v1/feeds", apiCfg.handleFeedsGetAll)

	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollowsAdd))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handleFeedFollowsDelete))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollowsGetAllByUser))

	srv := &http.Server{
		Addr:    ":" + env.Port,
		Handler: mux,
	}

	log.Println("Serving on port: ", env.Port)
	log.Fatal(srv.ListenAndServe())
}

func (cfg *apiConfig) fetchDataFromFeed(url string) (parsedXML, error) {
	resp, err := http.Get(url)
	if err != nil {
		return parsedXML{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return parsedXML{}, errors.New("Status for GET at URL not 200 (OK)")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return parsedXML{}, err
	}

	result := parsedXML{}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return parsedXML{}, err
	}

	return result, err
}

func (cfg *apiConfig) FetchFeeds() {
	ticker := time.NewTimer(cfg.fetch_interval)
	defer ticker.Stop()

	for {
		// TODO: Concurrently fetch cfg.fetch_count feeds - use sync.Waitgroup
	}
}
