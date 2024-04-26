package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/handlers/user"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Serve frontend files
	frontend := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", frontend)

	// API endpoints
	http.HandleFunc("/ping", handlers.HealthCheck)
	http.HandleFunc("/users/create", user.CreateUser)
	http.HandleFunc("/users/login", user.LoginUser)
	http.HandleFunc("/users/delete", user.DeleteUser)

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Get database and collections
	db := client.Database("syncscribe")
	usersCollection := db.Collection("users")
	//notesCollection := db.Collection("notes")
	//foldersCollection := db.Collection("folders")

	// Pass the MongoDB collections to the handlers
	handlers.SetCollections(usersCollection, nil, nil)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(http.DefaultServeMux)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
