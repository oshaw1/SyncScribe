package note

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	notesCollection := handlers.GetNotesCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve notes for the given user ID
	notesCursor, err := notesCollection.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer notesCursor.Close(ctx)

	var notes []models.Note
	for notesCursor.Next(ctx) {
		var note models.Note
		if err := notesCursor.Decode(&note); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	json.NewEncoder(w).Encode(notes)
}
