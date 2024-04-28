package user

import (
	"SyncScribe/backend/handlers"
	"SyncScribe/backend/handlers/response"
	"SyncScribe/backend/models"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usersCollection *mongo.Collection

func Init(mongoUsersCollection *mongo.Collection) {
	usersCollection = mongoUsersCollection
}

func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.SendErrorResponse(responseWriter, http.StatusBadRequest, "Error decoding request body", err)
		return
	}

	if user.Username == "" || user.Password == "" {
		response.SendErrorResponse(responseWriter, http.StatusBadRequest, "Missing required fields", nil)
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.SendErrorResponse(responseWriter, http.StatusInternalServerError, "Error hashing password", err)
		return
	}
	user.Password = string(hashedPassword)

	user.Notes = []string{}
	user.Allowed = false

	result, err := handlers.GetUsersCollection().InsertOne(context.Background(), user)
	if err != nil {
		response.SendErrorResponse(responseWriter, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	userID := result.InsertedID.(primitive.ObjectID).Hex()
	response.SendSuccessResponse(responseWriter, "User created successfully", map[string]interface{}{"userID": userID})
}
