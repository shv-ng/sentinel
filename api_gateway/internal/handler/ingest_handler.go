package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/service"
	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/utils"
)

type IngestHandler interface {
	HandleIngest(w http.ResponseWriter, r *http.Request)
}
type ingestHandler struct {
	service service.IngestorService
}

func NewIngestHandler(s service.IngestorService) IngestHandler {
	return &ingestHandler{s}
}

func (h *ingestHandler) HandleIngest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || !utils.IsValidJSON(body) {
		utils.ErrorJSON(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	res, err := h.service.SendLogParser(string(body))
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}
	if !res.Success {
		utils.ErrorJSON(w, http.StatusBadRequest, res.Message)
		return
	}

	utils.WriteJSON(w, http.StatusOK, res)
}
