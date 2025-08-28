package logs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ShivangSrivastava/sentinel/internal/logformat"
)

func ConvertFieldValue(value string, fieldType string, dateFormat string) (any, error) {
	switch fieldType {
	case "string":
		return value, nil
	case "int":
		return strconv.Atoi(value)
	case "float":
		return strconv.ParseFloat(value, 64)
	case "datetime":
		if dateFormat != "" {
			return time.Parse(dateFormat, value)
		}
		formats := []string{
			time.RFC3339,
			"2006-01-02 15:04:05",
			"02/Jan/2006:15:04:05 -0700",
		}
		for _, format := range formats {
			if t, err := time.Parse(format, value); err == nil {
				return t, nil
			}
		}
		return nil, fmt.Errorf("unable to parse datetime: %s", value)
	case "bool":
		return strconv.ParseBool(value)
	default:
		return value, nil
	}
}
func ValidateFields(parsedData map[string]any, fields []logformat.LogFormatField) error {
	for _, field := range fields {
		semanticName := field.SemanticName
		if semanticName == nil {
			semanticName = &field.RawName
		}

		value, exists := parsedData[*semanticName]
		if !exists && field.Required {
			return fmt.Errorf("required field missing: %s", *semanticName)
		}

		if exists && value == nil && field.Required {
			return fmt.Errorf("required field is null: %s", *semanticName)
		}
	}
	return nil
}
