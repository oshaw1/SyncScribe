package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oshaw1/SyncScribe/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection
var notesCollection *mongo.Collection
var foldersCollection *mongo.Collection

func SetCollections(users, notes, folders *mongo.Collection) {
	usersCollection = users
	notesCollection = notes
	foldersCollection = folders
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Notes = []string{}
	user.Allowed = false

	result, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	response := map[string]string{"userID": userID}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
