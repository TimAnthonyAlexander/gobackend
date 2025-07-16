package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/gobackend/internal/models"
)

// DateTimeHandler returns the current date and time
func DateTimeHandler(w http.ResponseWriter, r *http.Request) {
	response := models.NewDateTimeResponse()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
} 