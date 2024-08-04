package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	addr := "localhost:" + os.Getenv("PORT")
	godotenv.Load(".env")
	mux := http.NewServeMux()
	srv := &http.Server{Handler: mux, Addr: addr}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
