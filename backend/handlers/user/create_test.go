package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"SyncScribe/backend/handlers/user"
	"SyncScribe/backend/models"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser_Success(t *testing.T) {
	client, usersCollection := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	userReq := models.User{
		Username: "testuser",
		Password: "password",
		Notes:    []string{},
		Allowed:  false,
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

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	userID, ok := response["userID"].(string)
	require.True(t, ok)
	require.NotEmpty(t, userID)

	userIDObj, err := primitive.ObjectIDFromHex(userID)
	require.NoError(t, err)

	// Retrieve the created user from the database
	var createdUser models.User
	err = usersCollection.FindOne(context.Background(), primitive.M{"_id": userIDObj}).Decode(&createdUser)
	require.NoError(t, err)

	require.Equal(t, userReq.Username, createdUser.Username)
	require.Equal(t, userReq.Password, createdUser.Password)
	require.Equal(t, userReq.Notes, createdUser.Notes)
	require.Equal(t, userReq.Allowed, createdUser.Allowed)

	// Clean up the created user
	_, err = usersCollection.DeleteOne(context.Background(), primitive.M{"_id": userIDObj})
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
