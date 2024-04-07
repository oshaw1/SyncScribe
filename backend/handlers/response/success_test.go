package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oshaw1/SyncScribe/backend/handlers/response"
	"github.com/stretchr/testify/assert"
)

func TestSendSuccessResponse(t *testing.T) {
	rr := httptest.NewRecorder()

	response.SendSuccessResponse(rr, "Success")

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"message": "Success"}`, rr.Body.String())
}
