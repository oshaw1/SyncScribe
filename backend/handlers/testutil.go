package handlers

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupTestDatabase(t *testing.T) (*mongo.Client, *mongo.Collection, *mongo.Collection, *mongo.Collection) {
	uri := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(uri)
	if err != nil {
		t.Fatalf("Error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}

	db := client.Database("Test")
	usersCollection := db.Collection("users")
	notesCollection := db.Collection("notes")
	foldersCollection := db.Collection("folders")

	SetCollections(usersCollection, notesCollection, foldersCollection)

	return client, usersCollection, notesCollection, foldersCollection
}
