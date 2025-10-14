package logformat

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/shv-ng/sentinel/pkg/utils"
)

type LogFormatHandler interface {
	CreateLogFormat(w http.ResponseWriter, r *http.Request)
	GetFormatByName(w http.ResponseWriter, r *http.Request)
	GetAllFormats(w http.ResponseWriter, r *http.Request)
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
		utils.ErrorJSON(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = h.service.CreateLogFormat(string(body))
	if err != nil {
		utils.ErrorJSON(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, map[string]string{"message": "Log format successfully created."}, http.StatusOK)
}

func (h *handler) GetFormatByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		utils.ErrorJSON(w, "name parameter is required", http.StatusBadRequest)
		return
	}
	parser, fields, err := h.service.GetFormatByName(name)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorJSON(w, "parser not found", http.StatusNotFound)
			return
		}
		utils.ErrorJSON(w, "internal server error", http.StatusInternalServerError)
		return
	}
	res := struct {
		Parser *LogFormatParser `json:"parser"`
		Fields []LogFormatField `json:"fields"`
	}{
		Parser: parser,
		Fields: fields,
	}
	utils.WriteJSON(w, res, http.StatusOK)
}

func (h *handler) GetAllFormats(w http.ResponseWriter, r *http.Request) {
	parsers, err := h.service.GetAllFormats()
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorJSON(w, "parser not found", http.StatusNotFound)
			return
		}
		utils.ErrorJSON(w, "internal server error", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, parsers, http.StatusOK)
}
