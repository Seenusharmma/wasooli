package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"wasooli/backend/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Disconnect()

	// Set up HTTP routes
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Check database connection
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Database connection failed: %v", err)
			return
		}
		fmt.Fprintf(w, "OK")
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
