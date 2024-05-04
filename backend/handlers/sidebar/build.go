package sidebar

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

func BuildSidebarStructure(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		UserID string `json:"userId"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userID := requestData.UserID

	notes, err := fetchNotes(r.Context(), userID)
	if err != nil {
		log.Printf("Error fetching notes: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Retrieved notes: %+v", notes)

	folders, err := fetchFolders(r.Context(), userID)
	if err != nil {
		log.Printf("Error fetching folders: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Retrieved folders: %+v", folders)

	if len(notes) == 0 && len(folders) == 0 {
		log.Printf("No notes or folders found for userID: %s", userID)
		http.Error(w, "No notes or folders found for the provided userID", http.StatusNotFound)
		return
	}

	sidebarStructure := buildStructure(notes, folders)
	log.Printf("Sidebar structure: %+v", sidebarStructure)

	sidebarStructureBytes, err := json.Marshal(sidebarStructure)
	if err != nil {
		log.Printf("Error marshaling sidebar structure: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(sidebarStructureBytes)
}

func fetchNotes(ctx context.Context, userID string) ([]models.Note, error) {
	notesCollection := handlers.GetNotesCollection()
	notesCursor, err := notesCollection.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}
	defer notesCursor.Close(ctx)

	var notes []models.Note
	for notesCursor.Next(ctx) {
		var note models.Note
		if err := notesCursor.Decode(&note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func fetchFolders(ctx context.Context, userID string) ([]models.Folder, error) {
	foldersCollection := handlers.GetFoldersCollection()
	foldersCursor, err := foldersCollection.Find(ctx, bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}
	defer foldersCursor.Close(ctx)

	var folders []models.Folder
	for foldersCursor.Next(ctx) {
		var folder models.Folder
		if err := foldersCursor.Decode(&folder); err != nil {
			return nil, err
		}
		folders = append(folders, folder)
	}

	return folders, nil
}

func buildStructure(notes []models.Note, folders []models.Folder) map[string]interface{} {
	sidebarStructure := make(map[string]interface{})
	folderMap := make(map[string]models.Folder)
	noteMap := make(map[string]models.Note)

	for _, folder := range folders {
		folderMap[folder.ID.Hex()] = folder
	}

	for _, note := range notes {
		noteMap[note.ID.Hex()] = note
	}

	// Assign notes to their respective folders
	for _, note := range notes {
		if note.FolderID != "" {
			folder, ok := folderMap[note.FolderID]
			if ok {
				folder.NoteIDs = append(folder.NoteIDs, note.ID.Hex())
				folderMap[note.FolderID] = folder
			}
		} else {
			// Note doesn't belong to any folder, add it to the root level
			sidebarStructure[note.ID.Hex()] = map[string]interface{}{
				"id":   note.ID.Hex(),
				"name": note.Title,
				"type": "note",
			}
		}
	}

	// Build the folder hierarchy recursively
	var buildFolderHierarchy func(folderID string) interface{}
	buildFolderHierarchy = func(folderID string) interface{} {
		folder, ok := folderMap[folderID]
		if !ok {
			return nil
		}

		folderItem := map[string]interface{}{
			"id":       folder.ID.Hex(),
			"name":     folder.Name,
			"type":     "folder",
			"children": []interface{}{},
		}

		// Add child folders recursively
		for _, childFolderID := range folder.ChildFolderIDs {
			childFolderItem := buildFolderHierarchy(childFolderID)
			if childFolderItem != nil {
				folderItem["children"] = append(folderItem["children"].([]interface{}), childFolderItem)
			}
		}

		// Add notes
		for _, noteID := range folder.NoteIDs {
			if note, ok := noteMap[noteID]; ok {
				noteItem := map[string]interface{}{
					"id":   note.ID.Hex(),
					"name": note.Title,
					"type": "note",
				}
				folderItem["children"] = append(folderItem["children"].([]interface{}), noteItem)
			}
		}

		return folderItem
	}

	// Build the folder hierarchy starting from the root level folders
	for _, folder := range folderMap {
		if folder.ParentFolderID == "" {
			folderItem := buildFolderHierarchy(folder.ID.Hex())
			if folderItem != nil {
				sidebarStructure[folder.ID.Hex()] = folderItem
			}
		}
	}

	return sidebarStructure
}
