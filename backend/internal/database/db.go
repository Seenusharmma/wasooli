package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// Connect initializes a connection to MongoDB
func Connect() (*Database, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGODB_URI not set in environment variables")
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		dbName = "wasooli" // Default database name
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create MongoDB client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	log.Println("Successfully connected to MongoDB")

	return &Database{
		Client: client,
		DB:     client.Database(dbName),
	}, nil
}

// Disconnect closes the database connection
func (db *Database) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := db.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	log.Println("Disconnected from MongoDB")
	return nil
}

// Ping checks if the database connection is alive
func (db *Database) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return db.Client.Ping(ctx, readpref.Primary())
}
