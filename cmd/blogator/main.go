package main

// The Postgresql driver is needed for its side-effects.
import (
	"log"
	"os"

	"github.com/benskia/Blogator/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	dbURL := os.Getenv("CONN")
	if dbURL == "" {
		log.Fatal("CONN environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	api.StartBlogator(api.EnvVars{
		DbURL: dbURL,
		Port:  port,
	})
}
