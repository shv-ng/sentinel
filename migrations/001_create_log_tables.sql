CREATE TABLE IF NOT EXISTS log_parsers (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL UNIQUE,
  is_json BOOLEAN NOT NULL,
  regex_pattern TEXT,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS log_fields (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  parser_id UUID NOT NULL REFERENCES log_parsers(id) ON DELETE CASCADE,
  raw_name TEXT NOT NULL,
  semantic_name TEXT,
  type TEXT NOT NULL,
  datetime_format TEXT,
  enum_value TEXT,
  required BOOLEAN DEFAULT FALSE,
  description TEXT
);
