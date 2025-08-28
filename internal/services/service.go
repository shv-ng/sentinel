package logformat

import (
	"encoding/json"

	"github.com/ShivangSrivastava/sentinel/internal/shared"
)

type LogFormatService interface {
	CreateLogFormat(jsonData string) error
	GetFormatByName(name string) (*shared.LogFormatParser, []shared.LogFormatField, error)
	GetAllFormats() ([]shared.LogFormatParser, error)
}
type service struct {
	repo LogFormatRepo
}

func NewService(repo LogFormatRepo) LogFormatService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateLogFormat(jsonData string) error {
	var p struct {
		Name         string                  `json:"name"`
		IsJSON       bool                    `json:"is_json"`
		RegexPattern *string                 `json:"regex_pattern"`
		Fields       []shared.LogFormatField `json:"fields"`
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
	id, err := s.repo.CreateFormatParser(p.Name, p.IsJSON, regexPattern)
	if err != nil {
		return err
	}
	for _, field := range p.Fields {
		field.ParserID = id
		err := s.repo.CreateFormatField(field)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) GetFormatByName(name string) (*shared.LogFormatParser, []shared.LogFormatField, error) {
	return s.repo.GetByFormatName(name)
}

func (s *service) GetAllFormats() ([]shared.LogFormatParser, error) {
	return s.repo.GetAllFormats()
}
