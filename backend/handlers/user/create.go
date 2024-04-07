package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/oshaw1/SyncScribe/backend/handlers"
	"github.com/oshaw1/SyncScribe/backend/handlers/response"
	"github.com/oshaw1/SyncScribe/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection

func Init(uc *mongo.Collection) {
	usersCollection = uc
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	// Check if required fields are missing
	if user.Username == "" || user.Password == "" {
		response.SendErrorResponse(w, http.StatusBadRequest, "Missing required fields", nil)
		return
	}

	user.Notes = []string{}
	user.Allowed = false

	result, err := handlers.GetUsersCollection().InsertOne(context.Background(), user)
	if err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	response.SendSuccessResponse(w, "User created successfully", map[string]string{"userID": userID})
}
