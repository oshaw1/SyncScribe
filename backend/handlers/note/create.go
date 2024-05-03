package note

import (
	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note.ID = primitive.NewObjectID()
	note.CreatedAt = time.Now().Format(time.RFC3339)
	note.UpdatedAt = note.CreatedAt

	notesCollection := handlers.GetNotesCollection()
	_, err = notesCollection.InsertOne(context.Background(), note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}
