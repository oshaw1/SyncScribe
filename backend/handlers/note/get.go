package note

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNotes(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	notesCollection := handlers.GetNotesCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	noteID := r.URL.Query().Get("noteId")
	fmt.Println("Note ID:", noteID)

	objectID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	notesCollection := handlers.GetNotesCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve the note with the given note ID
	var note models.Note
	err = notesCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&note)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Note not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Log success request with the note title
	fmt.Printf("Successfully retrieved note: %s\n", note.Title)

	json.NewEncoder(w).Encode(note)
}
