package utils

import (
	"encoding/json"
	"net/http"
)

// DecodeJSON reads JSON into a target struct
func DecodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		ErrorJSON(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return false
	}
	return true
}

// WriteJSON sends a JSON response
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ErrorJSON sends a standard error message as JSON
func ErrorJSON(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}
