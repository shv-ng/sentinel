package logformat

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ShivangSrivastava/sentinel/pkg/utils"
)

type LogFormatHandler interface {
	CreateLogFormat(w http.ResponseWriter, r *http.Request)
}
type handler struct {
	service LogFormatService
}

func NewHandler(service LogFormatService) LogFormatHandler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateLogFormat(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || !utils.IsValidJSON(body) {
		utils.ErrorJSON(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	err = h.service.CreateLogFormat(string(body))
	if err != nil {
		utils.ErrorJSON(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Log format successfully created."})
}
