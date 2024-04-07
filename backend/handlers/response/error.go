package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	response := map[string]string{"message": message}
	if err != nil {
		response["error"] = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
