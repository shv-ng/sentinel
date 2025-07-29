package repo

import (
	"database/sql"

	"github.com/ShivangSrivastava/sentinel/ingestor/internal/models"
	"github.com/google/uuid"
)

type LogParserRepo interface {
	CreateLogParser(name string, is_json bool, regex_pattern string) (uuid.UUID, error)
	CreateLogField(arg models.LogField) error
}

type logParserRepo struct {
	db *sql.DB
}

func NewLogParserRepo(db *sql.DB) LogParserRepo {
	return &logParserRepo{
		db: db,
	}
}

func (l *logParserRepo) CreateLogParser(name string, is_json bool, regex_pattern string) (uuid.UUID, error) {
	q := `INSERT INTO log_parsers (name, is_json, regex_pattern)
VALUES ($1, $2, $3)
RETURNING id;`
	var id uuid.UUID
	err := l.db.QueryRow(q, name, is_json, regex_pattern).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (l *logParserRepo) CreateLogField(arg models.LogField) error {
	q := `INSERT INTO log_fields (
  parser_id, raw_name, semantic_name, type, datetime_format,
  enum_value, required, description
)
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
);`
	_, err := l.db.Exec(
		q, arg.ParserID, arg.RawName, arg.SemanticName, arg.Type,
		arg.DatetimeFormat, arg.EnumValue, arg.Required, arg.Description,
	)

	if err != nil {
		return err
	}
	return nil
}
