package folder

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetFolders(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")

	foldersCollection := handlers.GetFoldersCollection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve folders for the given user ID
	foldersCursor, err := foldersCollection.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer foldersCursor.Close(ctx)

	var folders []models.Folder
	for foldersCursor.Next(ctx) {
		var folder models.Folder
		if err := foldersCursor.Decode(&folder); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		folders = append(folders, folder)
	}

	json.NewEncoder(w).Encode(folders)
}
