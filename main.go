package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("GO Server")

	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access the PORT environment variable
	portstring := os.Getenv("PORT")
	if portstring == "" {
		log.Fatal("PORT Not Found!")
	}

	// Initialize router
	router := chi.NewRouter()

	// Use CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/healthZ", handlerReadniess)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	// Start the server
	fmt.Println("Starting server on port", portstring)
	log.Fatal(http.ListenAndServe(":"+portstring, router))
}
