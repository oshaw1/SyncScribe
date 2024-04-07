package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oshaw1/SyncScribe/backend/handlers/user"
	"github.com/oshaw1/SyncScribe/backend/models"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser_Success(t *testing.T) {
	client, usersCollection := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	userReq := models.User{
		Username: "testuser",
		Password: "password",
	}
	body, err := json.Marshal(userReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "users/create", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.CreateUser)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var response struct {
		Message string `json:"message"`
		Data    struct {
			UserID string `json:"userID"`
		} `json:"data"`
	}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	userID, err := primitive.ObjectIDFromHex(response.Data.UserID)
	require.NoError(t, err)

	_, err = usersCollection.DeleteOne(context.Background(), primitive.M{"_id": userID})
	require.NoError(t, err)
}

func TestCreateUser_InvalidRequestBody(t *testing.T) {
	client, _ := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	req, err := http.NewRequest("POST", "users/create", bytes.NewBufferString("invalid json"))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.CreateUser)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateUser_MissingFields(t *testing.T) {
	client, _ := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	userReq := models.User{
		Username: "testuser",
	}
	body, err := json.Marshal(userReq)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "users/create", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.CreateUser)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code)
}
