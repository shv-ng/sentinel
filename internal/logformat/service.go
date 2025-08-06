package logformat

import "encoding/json"

type LogFormatService interface {
	CreateLogFormat(jsonData string) error
	GetFormatByName(name string) (*LogFormatParser, []LogFormatField, error)
	GetAllFormats() ([]LogFormatParser, error)
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
		Name         string           `json:"name"`
		IsJson       bool             `json:"is_json"`
		RegexPattern *string          `json:"regex_pattern"`
		Fields       []LogFormatField `json:"fields"`
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
	id, err := s.repo.CreateFormatParser(p.Name, p.IsJson, regexPattern)
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

func (s *service) GetFormatByName(name string) (*LogFormatParser, []LogFormatField, error) {
	return s.repo.GetByFormatName(name)
}

func (s *service) GetAllFormats() ([]LogFormatParser, error) {
	return s.repo.GetAllFormats()
}
