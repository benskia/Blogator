package main

// The Postgresql driver is needed for its side-effects.
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/benskia/Blogator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	dbURL := os.Getenv("CONN")
	if dbURL == "" {
		log.Fatal("CONN environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}
	dbQueries := database.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handleReadiness)
	mux.HandleFunc("GET /v1/err", handleError)
	mux.HandleFunc("POST /v1/users", createUser)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Printf("Serving on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
