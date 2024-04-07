package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/oshaw1/SyncScribe/backend/handlers"
	"github.com/oshaw1/SyncScribe/backend/handlers/response"
	"github.com/oshaw1/SyncScribe/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Error decoding login data", err)
		return
	}

	user, err := findUserByCredentials(loginData.Username, loginData.Password)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			response.SendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials", nil)
		} else {
			response.SendErrorResponse(w, http.StatusInternalServerError, "Error finding user", err)
		}
		return
	}

	if !user.Allowed {
		response.SendErrorResponse(w, http.StatusForbidden, "Access denied", nil)
		return
	}

	response.SendSuccessResponse(w, "Login successful")
}

func findUserByCredentials(username, password string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username, "password": password}
	err := handlers.GetUsersCollection().FindOne(context.Background(), filter).Decode(&user)
	return user, err
}
