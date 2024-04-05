package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/oshaw1/SyncScribe/backend/models"
	"go.mongodb.org/mongo-driver/bson"
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding login data: %v", err)
		return
	}

	var user models.User
	filter := bson.M{"username": loginData.Username, "password": loginData.Password}
	err = usersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			response := map[string]string{"message": "Invalid credentials"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("Error encoding response: %v", err)
			}
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error finding user: %v", err)
		return
	}

	if !user.Allowed {
		response := map[string]string{"message": "Access denied"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
		return
	}

	response := map[string]string{"message": "Login successful"}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
}
