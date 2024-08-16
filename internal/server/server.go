package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/benskia/Blogator/internal/database"
	_ "github.com/lib/pq" // The psql driver is needed for its side-effects.
)

type EnvVars struct {
	DbURL string
	Port  string
}

type apiConfig struct {
	DB *database.Queries
}

func StartBlogator(env EnvVars) {
	db, err := sql.Open("postgres", env.DbURL)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
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
