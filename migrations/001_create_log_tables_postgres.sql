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

CREATE TABLE IF NOT EXISTS logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    parser_id UUID NOT NULL REFERENCES log_parsers(id),
    source_file TEXT NOT NULL,
    parsed_data JSONB NOT NULL,
    raw_log TEXT NOT NULL,
    ingestion_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_logs_timestamp ON logs(timestamp);
CREATE INDEX IF NOT EXISTS idx_logs_parser_id ON logs(parser_id);
CREATE INDEX IF NOT EXISTS idx_logs_ingestion_time ON logs(ingestion_time);

-- GIN index for JSONB queries
CREATE INDEX IF NOT EXISTS idx_logs_parsed_data ON logs USING GIN(parsed_data);
