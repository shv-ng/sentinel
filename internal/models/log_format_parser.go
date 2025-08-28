package models

import (
	"time"

	"github.com/google/uuid"
)

type LogFormatParser struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	IsJSON       bool      `json:"is_json"`
	RegexPattern *string   `json:"regex_pattern"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
