package user

import (
	"SyncScribe/backend/handlers"
	"SyncScribe/backend/handlers/response"
	"SyncScribe/backend/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	filter := bson.M{"username": username}
	err := handlers.GetUsersCollection().FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		fmt.Printf("Error finding user: %v\n", err)
		return user, err
	}

	fmt.Printf("Retrieved user: %+v\n", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Printf("Password comparison failed: %v\n", err)
		return user, mongo.ErrNoDocuments
	}

	return user, nil
}

func GenerateJWTToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
	})
	return token.SignedString(JWTSecret)
}
