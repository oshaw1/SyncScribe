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
		sendErrorResponse(w, http.StatusBadRequest, "Error decoding login data", err)
		return
	}

	user, err := findUserByCredentials(loginData.Username, loginData.Password)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			sendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials", nil)
		} else {
			sendErrorResponse(w, http.StatusInternalServerError, "Error finding user", err)
		}
		return
	}

	if !user.Allowed {
		sendErrorResponse(w, http.StatusForbidden, "Access denied", nil)
		return
	}

	sendSuccessResponse(w, "Login successful")
}

func findUserByCredentials(username, password string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username, "password": password}
	err := usersCollection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
	if err != nil {
		log.Printf("Error: %v", err)
	}
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}
