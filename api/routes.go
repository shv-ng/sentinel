package api

import (
	"net/http"

	"github.com/ShivangSrivastava/sentinel/internal/logformat"
	"github.com/ShivangSrivastava/sentinel/pkg/middleware"
)

func NewRouter(logformatHandler logformat.LogFormatHandler) http.Handler {
	mux := http.NewServeMux()

	// Register API routes
	registerLogformatRoutes(mux, logformatHandler)

	handler := middleware.LoggingMiddleware(mux)

	return http.StripPrefix("/v1", handler)
}

func registerLogformatRoutes(mux *http.ServeMux, h logformat.LogFormatHandler) {
	mux.HandleFunc("POST /log-formats", h.CreateLogFormat)
	mux.HandleFunc("GET /log-formats/{name}", h.GetFormatByName)

}
