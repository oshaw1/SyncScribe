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

func TestLoginUser_SuccessfulLogin(t *testing.T) {
	client, usersCollection := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "password",
		Allowed:  true,
	}
	_, err := usersCollection.InsertOne(context.Background(), testUser)
	require.NoError(t, err)

	// Test successful login
	loginData := map[string]string{
		"username": "testuser",
		"password": "password",
	}
	body, err := json.Marshal(loginData)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.LoginUser)
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, "Login successful", response["message"])

	// Clean up the test user
	_, err = usersCollection.DeleteOne(context.Background(), primitive.M{"username": "testuser"})
	require.NoError(t, err)
}

func TestLoginUser_InvalidCredentials(t *testing.T) {
	client, usersCollection := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "password",
		Allowed:  true,
	}
	_, err := usersCollection.InsertOne(context.Background(), testUser)
	require.NoError(t, err)

	// Test invalid credentials
	loginData := map[string]string{
		"username": "testuser",
		"password": "wrongpassword",
	}
	body, err := json.Marshal(loginData)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.LoginUser)
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusUnauthorized, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, "Invalid credentials", response["message"])

	// Clean up the test user
	_, err = usersCollection.DeleteOne(context.Background(), primitive.M{"username": "testuser"})
	require.NoError(t, err)
}

func TestLoginUser_AccessDenied(t *testing.T) {
	client, usersCollection := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "password",
		Allowed:  false,
	}
	_, err := usersCollection.InsertOne(context.Background(), testUser)
	require.NoError(t, err)

	// Test access denied
	loginData := map[string]string{
		"username": "testuser",
		"password": "password",
	}
	body, err := json.Marshal(loginData)
	require.NoError(t, err)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.LoginUser)
	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusForbidden, rr.Code)
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Equal(t, "Access denied", response["message"])

	// Clean up the test user
	_, err = usersCollection.DeleteOne(context.Background(), primitive.M{"username": "testuser"})
	require.NoError(t, err)
}
