package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"SyncScribe/backend/handlers"
	"SyncScribe/backend/handlers/response"
	"SyncScribe/backend/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var JWTSecret = []byte("9jlvzXsJzf+QuNSlqTHrAZ1FaAbGVEMyqqaCeHkoKwg=")

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
		fmt.Println("Access denied for user:", user.Username)
		response.SendErrorResponse(w, http.StatusForbidden, "Access denied", nil)
		return
	}

	token, err := GenerateJWTToken(user.ID.Hex())
	if err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, "Error generating token", err)
		return
	}

	response.SendSuccessResponse(w, "Login successful", map[string]interface{}{"token": token})
}

func findUserByCredentials(username, password string) (models.User, error) {
	var user models.User
	filter := bson.M{"username": username, "password": password}
	err := handlers.GetUsersCollection().FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func GenerateJWTToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
	})
	return token.SignedString(JWTSecret)
}
