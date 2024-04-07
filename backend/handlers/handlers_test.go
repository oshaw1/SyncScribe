package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oshaw1/SyncScribe/backend/handlers"
	"github.com/oshaw1/SyncScribe/backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthCheck)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "pong", rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	// Create a mock MongoDB collection
	uri := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(uri)
	if err != nil {
		t.Fatalf("Error creating MongoDB client: %v", err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}

	usersCollection := client.Database("users").Collection("users")
	handlers.SetCollections(usersCollection, nil, nil)

	// Create a new user request
	user := models.User{
		Username: "testuser",
		Password: "password",
	}
	body, err := json.Marshal(user)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "users/create", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)

	userID, ok := response["userID"]
	require.True(t, ok, "userID key not found in response")

	objectID, err := primitive.ObjectIDFromHex(userID)
	require.NoError(t, err)

	_, err = usersCollection.DeleteOne(context.Background(), primitive.M{"_id": objectID})
	require.NoError(t, err)
}

func TestLoginUser_SuccessfulLogin(t *testing.T) {
	// Create a mock MongoDB collection
	uri := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(uri)
	if err != nil {
		t.Fatalf("Error creating MongoDB client: %v", err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}
	usersCollection := client.Database("users").Collection("users")
	handlers.SetCollections(usersCollection, nil, nil)

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "password",
		Allowed:  true,
	}
	_, err = usersCollection.InsertOne(context.Background(), testUser)
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
	handler := http.HandlerFunc(handlers.LoginUser)
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
	// Create a mock MongoDB collection
	uri := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(uri)
	if err != nil {
		t.Fatalf("Error creating MongoDB client: %v", err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}
	usersCollection := client.Database("users").Collection("users")
	handlers.SetCollections(usersCollection, nil, nil)

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "password",
		Allowed:  true,
	}
	_, err = usersCollection.InsertOne(context.Background(), testUser)
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
	handler := http.HandlerFunc(handlers.LoginUser)
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
	// Create a mock MongoDB collection
	uri := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(uri)
	if err != nil {
		t.Fatalf("Error creating MongoDB client: %v", err)
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		t.Fatalf("Error connecting to MongoDB: %v", err)
	}
	usersCollection := client.Database("users").Collection("users")
	handlers.SetCollections(usersCollection, nil, nil)

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "password",
		Allowed:  false,
	}
	_, err = usersCollection.InsertOne(context.Background(), testUser)
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
	handler := http.HandlerFunc(handlers.LoginUser)
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
