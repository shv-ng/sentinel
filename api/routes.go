package api

import (
	"net/http"

	"github.com/ShivangSrivastava/sentinel/internal/logformat"
)

func NewRouter(logformatHandler logformat.LogFormatHandler) http.Handler {
	mux := http.NewServeMux()

	// Register API routes
	registerLogformatRoutes(mux, logformatHandler)

	return http.StripPrefix("/v1", mux)
}

func registerLogformatRoutes(mux *http.ServeMux, h logformat.LogFormatHandler) {
	mux.HandleFunc("POST /log-formats", h.CreateLogFormat)
}
