package service

import (
	"encoding/json"

	"github.com/ShivangSrivastava/sentinel/ingestor/internal/models"
	"github.com/ShivangSrivastava/sentinel/ingestor/internal/repo"
)

type LogParserService interface {
	CreateLogParser(jsonData string) error
}

type logParserService struct {
	repo repo.LogParserRepo
}

func NewLogParserService(repo repo.LogParserRepo) LogParserService {
	return &logParserService{repo: repo}
}

func (l *logParserService) CreateLogParser(jsonData string) error {
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
	id, err := l.repo.CreateLogParser(p.Name, p.IsJson, *p.RegexPattern)
	if err != nil {
		return err
	}
	for _, field := range p.Fields {
		field.ParserID = id
		err := l.repo.CreateLogField(field)
		if err != nil {
			return err
		}
	}
	return nil
}
