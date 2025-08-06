package logformat

import (
	"database/sql"

	"github.com/google/uuid"
)

type LogFormatRepo interface {
	CreateFormatParser(name string, is_json bool, regex_pattern string) (uuid.UUID, error)
	CreateFormatField(arg logFormatField) error
}
type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) LogFormatRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateFormatParser(name string, is_json bool, regex_pattern string) (uuid.UUID, error) {
	q := `INSERT INTO log_parsers (name, is_json, regex_pattern)
	VALUES ($1, $2, $3)
	RETURNING id;`
	var id uuid.UUID
	err := r.db.QueryRow(q, name, is_json, regex_pattern).Scan(&id)
	if err != nil {
		log.Printf("[DB ERROR] failed to intsert into log_parsers table '%s': %v", name, err)
		return id, err
	}
	log.Printf("[DB SUCCESS] Inserted parser: %s (id: %s)", name, id)
	return id, nil
}

func (r *repo) CreateFormatField(arg logFormatField) error {
	q := `INSERT INTO log_fields (
	parser_id, raw_name, semantic_name, type, datetime_format,
	enum_value, required, description
	)
	VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8
	);`
	_, err := r.db.Exec(
		q, arg.ParserID, arg.RawName, arg.SemanticName, arg.Type,
		arg.DatetimeFormat, arg.EnumValue, arg.Required, arg.Description,
	)

	if err != nil {
		log.Printf("[DB ERROR] failed to intsert into log_fields table '%s': %v", arg.RawName, err)
		return err
	}
	log.Printf("[DB SUCCESS] Inserted fields: %s", arg.RawName)
	return nil
}
