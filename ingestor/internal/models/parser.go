package models

import (
	"time"

	"github.com/google/uuid"
)

type LogField struct {
	ID             uuid.UUID `json:"id"`
	ParserID       uuid.UUID `json:"parser_id"`
	RawName        string    `json:"raw_name"`
	SemanticName   *string   `json:"semantic_name"`
	Type           string    `json:"type"`
	DatetimeFormat *string   `json:"datetime_format"`
	EnumValue      *string   `json:"enum_value"`
	Required       *string   `json:"required"`
	Description    *string   `json:"description"`
}

type LogParser struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	IsJson       bool      `json:"is_json"`
	RegexPattern *string   `json:"regex_pattern"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
