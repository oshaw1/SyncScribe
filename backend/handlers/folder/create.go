package folder

import (
	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFolder(w http.ResponseWriter, r *http.Request) {
	var folder models.Folder
	err := json.NewDecoder(r.Body).Decode(&folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	folder.ID = primitive.NewObjectID()

	foldersCollection := handlers.GetFoldersCollection()
	_, err = foldersCollection.InsertOne(context.Background(), folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(folder)
}
