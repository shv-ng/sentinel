package logformat

import (
	"database/sql"

	"github.com/google/uuid"
)

type LogFormatRepo interface {
	CreateFormatParser(name string, is_json bool, regex_pattern string) (uuid.UUID, error)
	CreateFormatField(arg LogFormatField) error
	GetByFormatName(name string) (*LogFormatParser, []LogFormatField, error)
	GetAllFormats() ([]LogFormatParser, error)
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
		return id, err
	}
	return id, nil
}

func (r *repo) CreateFormatField(arg LogFormatField) error {
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
		return err
	}
	return nil
}

func (r *repo) GetByFormatName(name string) (*LogFormatParser, []LogFormatField, error) {
	q := `SELECT id, name, is_json, regex_pattern, created_at, updated_at 
	FROM log_parsers WHERE name = $1`
	var parser LogFormatParser
	err := r.db.QueryRow(q, name).Scan(
		&parser.ID, &parser.Name, &parser.IsJson, &parser.RegexPattern,
		&parser.CreatedAt, &parser.UpdatedAt,
	)
	if err != nil {
		return nil, nil, err
	}

	// Get fields for this parser
	fields, err := r.getFieldsByParserID(parser.ID)
	if err != nil {
		return nil, nil, err
	}
	return &parser, fields, nil
}

func (r *repo) GetAllFormats() ([]LogFormatParser, error) {
	q := `SELECT id, name, is_json, regex_pattern, created_at, updated_at 
	FROM log_parsers ORDER BY created_at DESC`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parsers []LogFormatParser
	for rows.Next() {
		var parser LogFormatParser
		err := rows.Scan(
			&parser.ID, &parser.Name, &parser.IsJson, &parser.RegexPattern,
			&parser.CreatedAt, &parser.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		parsers = append(parsers, parser)
	}
	return parsers, nil
}

func (r *repo) getFieldsByParserID(parserID uuid.UUID) ([]LogFormatField, error) {
	q := `SELECT id, parser_id, raw_name, semantic_name, type, 
	datetime_format, enum_value, required, description
	FROM log_fields WHERE parser_id = $1 ORDER BY raw_name`

	rows, err := r.db.Query(q, parserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []LogFormatField
	for rows.Next() {
		var field LogFormatField
		err := rows.Scan(
			&field.ID, &field.ParserID, &field.RawName, &field.SemanticName,
			&field.Type, &field.DatetimeFormat, &field.EnumValue,
			&field.Required, &field.Description,
		)
		if err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	return fields, nil
}
