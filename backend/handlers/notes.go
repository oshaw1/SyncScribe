package handlers

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection
var notesCollection *mongo.Collection
var foldersCollection *mongo.Collection

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func HandleNotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SetCollections(users, notes, folders *mongo.Collection) {
	usersCollection = users
	notesCollection = notes
	foldersCollection = folders
}
