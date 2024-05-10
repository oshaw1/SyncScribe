package main

import (
	"SyncScribe/backend/handlers"
	"SyncScribe/backend/handlers/folder"
	"SyncScribe/backend/handlers/note"
	"SyncScribe/backend/handlers/sidebar"
	"SyncScribe/backend/handlers/user"
	"context"
	"log"
	"net/http"
	"os"

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
	http.HandleFunc("/api/sidebar/build", sidebar.BuildSidebarStructure)

	// User endpoints
	http.HandleFunc("/users/create", user.CreateUser)
	http.HandleFunc("/users/login", user.LoginUser)
	http.HandleFunc("/users/delete", user.DeleteUser)

	// Note endpoints
	http.HandleFunc("/notes/create", note.CreateNote)
	http.HandleFunc("/notes/getNotes", note.GetNotes)
	http.HandleFunc("/notes/getNote", note.GetNoteByID)

	// Folder endpoints
	http.HandleFunc("/folders/create", folder.CreateFolder)
	http.HandleFunc("/folders/getFolders", folder.GetFolders)

	// WebSocket endpoint
	http.HandleFunc("/ws/", handlers.HandleWebSocket)

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("syncscribe")
	usersCollection := db.Collection("users")
	notesCollection := db.Collection("notes")
	foldersCollection := db.Collection("folders")
	handlers.SetCollections(usersCollection, notesCollection, foldersCollection)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	handler := c.Handler(http.DefaultServeMux)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
