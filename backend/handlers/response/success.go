package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, message string, data ...interface{}) {
	response := map[string]interface{}{"message": message}
	if len(data) > 0 {
		response["data"] = data[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
