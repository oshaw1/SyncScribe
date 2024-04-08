package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, message string, data map[string]interface{}) {
	response := map[string]interface{}{
		"message": message,
	}

	if data != nil {
		for key, value := range data {
			response[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
