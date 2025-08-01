package service

import (
	"encoding/json"

	"github.com/ShivangSrivastava/sentinel/ingestor/internal/models"
	"github.com/ShivangSrivastava/sentinel/ingestor/internal/repo"
)

type LogParserService interface {
	CreateLogFormat(jsonData string) error
}

type logParserService struct {
	repo repo.LogParserRepo
}

func NewLogParserService(repo repo.LogParserRepo) LogParserService {
	return &logParserService{repo: repo}
}

func (s *logParserService) CreateLogFormat(jsonData string) error {
	var p struct {
		Name         string            `json:"name"`
		IsJson       bool              `json:"is_json"`
		RegexPattern *string           `json:"regex_pattern"`
		Fields       []models.LogField `json:"fields"`
	}

	err := json.Unmarshal([]byte(jsonData), &p)

	if err != nil {
		return err
	}
	var regexPattern string
	if p.RegexPattern != nil {
		regexPattern = *p.RegexPattern
	} else {
		regexPattern = ""
	}
	if p.Name == "" {
		return nil
	}
	id, err := s.repo.CreateLogParserFmt(p.Name, p.IsJson, regexPattern)
	if err != nil {
		return err
	}
	for _, field := range p.Fields {
		field.ParserID = id
		err := s.repo.CreateLogField(field)
		if err != nil {
			return err
		}
	}
	return nil
}
