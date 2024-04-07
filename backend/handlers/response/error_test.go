package response_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oshaw1/SyncScribe/backend/handlers/response"
	"github.com/stretchr/testify/assert"
)

func TestSendErrorResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	err := errors.New("test error")

	response.SendErrorResponse(rr, http.StatusInternalServerError, "Internal Server Error", err)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"error":"test error", "message":"Internal Server Error"}`, rr.Body.String())
}
