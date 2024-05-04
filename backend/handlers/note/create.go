package note

import (
	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Title    string   `json:"title"`
		Content  string   `json:"content"`
		Tags     []string `json:"tags"`
		FolderID string   `json:"folderId"`
		UserID   string   `json:"userID"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Received UserID: %s", requestData.UserID)

	note := models.Note{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
		Title:     requestData.Title,
		Content:   requestData.Content,
		Tags:      requestData.Tags,
		UserID:    []string{requestData.UserID},
		FolderID:  requestData.FolderID,
	}

	notesCollection := handlers.GetNotesCollection()
	_, err = notesCollection.InsertOne(context.Background(), note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}
