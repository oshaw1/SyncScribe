package user_test

import (
	"context"
	"testing"
	"time"

	"SyncScribe/backend/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDatabase(t *testing.T) (*mongo.Client, *mongo.Collection) {
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

	usersCollection := client.Database("users").Collection("users")
	handlers.SetCollections(usersCollection, nil, nil)

	return client, usersCollection
}
