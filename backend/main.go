package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/oshaw1/SyncScribe/backend/handlers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Pin      string             `bson:"pin"`
	Notes    []string           `bson:"notes"`
	Allowed  []string           `bson:"allowed"`
}

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	NoteID    string             `bson:"noteID"`
	CreatedAt string             `bson:"createdAt"`
	Content   string             `bson:"content"`
	Tags      []string           `bson:"tags"`
	Title     string             `bson:"title"`
	UpdatedAt string             `bson:"updatedAt"`
	UserID    string             `bson:"userID"`
	FolderID  string             `bson:"folderID"`
}

type Folder struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	UserID         string             `bson:"userID"`
	ParentFolderID string             `bson:"parentFolderID"`
	ChildFolderIDs []string           `bson:"childFolderIDs"`
	NoteIDs        []string           `bson:"noteIDs"`
}

func main() {
	// Serve frontend files
	frontend := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", frontend)

	// API endpoint
	http.HandleFunc("/api/notes", handlers.HandleNotes)
	http.HandleFunc("/ping", handlers.HealthCheck)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("notes_app")
	usersCollection := db.Collection("users")
	notesCollection := db.Collection("notes")
	foldersCollection := db.Collection("folders")

	// Pass the MongoDB collections to the handlers
	handlers.SetCollections(usersCollection, notesCollection, foldersCollection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
