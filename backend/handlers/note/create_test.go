package note_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/handlers/note"
	"SyncScribe/backend/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateNote(t *testing.T) {
	client, _, notesCollection, _ := handlers.SetupTestDatabase(t)
	defer client.Disconnect(context.Background())

	// Create a new HTTP request with a JSON payload
	noteData := models.Note{
		Title:    "Test Note",
		Content:  "This is a test note.",
		Tags:     []string{"test", "note"},
		UserID:   []string{"user123"},
		FolderID: "folder456",
	}
	payload, _ := json.Marshal(noteData)
	req, _ := http.NewRequest("POST", "/notes/create", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the CreateNote handler
	handler := http.HandlerFunc(note.CreateNote)
	handler.ServeHTTP(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Decode the response body
	var createdNote models.Note
	err := json.NewDecoder(rr.Body).Decode(&createdNote)
	require.NoError(t, err)

	// Assert the created note fields
	assert.NotEqual(t, primitive.NilObjectID, createdNote.ID)
	assert.Equal(t, noteData.Title, createdNote.Title)
	assert.Equal(t, noteData.Content, createdNote.Content)
	assert.Equal(t, noteData.Tags, createdNote.Tags)
	assert.Equal(t, noteData.UserID, createdNote.UserID)
	assert.Equal(t, noteData.FolderID, createdNote.FolderID)

	// Retrieve the created note from the database
	var retrievedNote models.Note
	err = notesCollection.FindOne(context.Background(), primitive.M{"_id": createdNote.ID}).Decode(&retrievedNote)
	require.NoError(t, err)

	// Assert the retrieved note fields
	assert.Equal(t, createdNote.ID, retrievedNote.ID)
	assert.Equal(t, createdNote.Title, retrievedNote.Title)
	assert.Equal(t, createdNote.Content, retrievedNote.Content)
	assert.Equal(t, createdNote.Tags, retrievedNote.Tags)
	assert.Equal(t, createdNote.UserID, retrievedNote.UserID)
	assert.Equal(t, createdNote.FolderID, retrievedNote.FolderID)

	_, err = notesCollection.DeleteOne(context.Background(), primitive.M{"_id": createdNote.ID})
	require.NoError(t, err)
}

func TestCreateNote_InvalidRequestBody(t *testing.T) {
	_, _, _, _ = handlers.SetupTestDatabase(t)

	// Create an invalid HTTP request with malformed JSON payload
	req, _ := http.NewRequest("POST", "/notes/create", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the CreateNote handler
	handler := http.HandlerFunc(note.CreateNote)
	handler.ServeHTTP(rr, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
