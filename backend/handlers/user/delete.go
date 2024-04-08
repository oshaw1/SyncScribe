package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/oshaw1/SyncScribe/backend/handlers"
	"github.com/oshaw1/SyncScribe/backend/handlers/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromToken(r)
	if err != nil {
		response.SendErrorResponse(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		response.SendErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	filter := bson.M{"_id": objectID}
	result, err := handlers.GetUsersCollection().DeleteOne(context.Background(), filter)
	if err != nil {
		response.SendErrorResponse(w, http.StatusInternalServerError, "Error deleting user", err)
		return
	}

	if result.DeletedCount == 0 {
		response.SendErrorResponse(w, http.StatusNotFound, "User not found", nil)
		return
	}

	response.SendSuccessResponse(w, "User deleted successfully", nil)
}

func getUserIDFromToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in token claims")
	}

	return userID, nil
}
