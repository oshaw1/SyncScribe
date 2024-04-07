package handlers_test

import (
	"testing"

	"github.com/oshaw1/SyncScribe/backend/handlers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestSetCollections(t *testing.T) {
	// Create mock collections
	usersCollection := &mongo.Collection{}
	notesCollection := &mongo.Collection{}
	foldersCollection := &mongo.Collection{}

	// Set the collections
	handlers.SetCollections(usersCollection, notesCollection, foldersCollection)

	// Assert that the collections are set correctly
	assert.Equal(t, usersCollection, handlers.GetUsersCollection())
	assert.Equal(t, notesCollection, handlers.GetNotesCollection())
	assert.Equal(t, foldersCollection, handlers.GetFoldersCollection())
}
