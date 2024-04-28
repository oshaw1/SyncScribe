package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"SyncScribe/backend/handlers/user"
	"SyncScribe/backend/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteUser(t *testing.T) {
	client, usersCollection := setupTestDatabase(t)
	defer client.Disconnect(context.Background())

	// Create a test user
	testUser := models.User{
		ID:       primitive.NewObjectID(),
		Username: "testuser",
		Password: "password",
		Notes:    []string{},
		Allowed:  true,
	}
	_, err := usersCollection.InsertOne(context.Background(), testUser)
	require.NoError(t, err)

	// Generate a JWT token for the test user
	token, err := user.GenerateJWTToken(testUser.ID.Hex())
	require.NoError(t, err)

	req, err := http.NewRequest("DELETE", "/users", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(user.DeleteUser)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "User deleted successfully", response["message"])

	// Check that the user is deleted from the database
	var deletedUser models.User
	err = usersCollection.FindOne(context.Background(), primitive.M{"_id": testUser.ID}).Decode(&deletedUser)
	assert.Equal(t, mongo.ErrNoDocuments, err)
}
