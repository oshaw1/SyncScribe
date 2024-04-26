package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"SyncScribe/backend/handlers/response"

	"github.com/stretchr/testify/assert"
)

func TestSendSuccessResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	data := map[string]interface{}{
		"token": "sample-token",
		"id":    "user-id",
	}
	response.SendSuccessResponse(rr, "Success", data)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{
		"message": "Success",
		"token": "sample-token",
		"id": "user-id"
	}`, rr.Body.String())
}
