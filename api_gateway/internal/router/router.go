package router

import (
	"net/http"

	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/handler"
)

func NewRouter(ingestHandler handler.IngestHandler) http.Handler {

	mux := http.NewServeMux()

	registerIngestRoutes(mux, ingestHandler)

	return http.StripPrefix("/v1", mux)
}
