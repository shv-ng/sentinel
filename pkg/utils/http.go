package utils

import (
	"encoding/json"
	"net/http"
)

// DecodeJSON reads JSON into a target struct
func DecodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		ErrorJSON(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

// WriteJSON sends a JSON response
func WriteJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ErrorJSON sends a standard error message as JSON
func ErrorJSON(w http.ResponseWriter, message string, status int) {
	WriteJSON(w, map[string]string{"error": message}, status)
}
