package logs

import (
	"reflect"
	"testing"
	"time"

	"github.com/ShivangSrivastava/sentinel/internal/logformat"
)

func TestConvertFieldValue(t *testing.T) {
	tests := []struct {
		name       string
		value      string
		fieldType  string
		dateFormat string
		expected   any
		expectErr  bool
	}{
		{"String conversion", "hello", "string", "", "hello", false},
		{"Integer conversion", "42", "int", "", 42, false},
		{"Float conversion", "3.14", "float", "", 3.14, false},
		{"Boolean conversion", "true", "bool", "", true, false},
		{"Invalid int", "abc", "int", "", nil, true},
		{"Custom datetime format", "2025-08-06 14:00:00", "datetime", "2006-01-02 15:04:05", time.Date(2025, 8, 6, 14, 0, 0, 0, time.UTC), false},
		{"Fallback datetime format", "2025-08-06T14:00:00Z", "datetime", "", time.Date(2025, 8, 6, 14, 0, 0, 0, time.UTC), false},
		{"Invalid datetime", "not-a-date", "datetime", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ConvertFieldValue(tt.value, tt.fieldType, tt.dateFormat)

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else if tt.fieldType == "datetime" {
					if !result.(time.Time).Equal(tt.expected.(time.Time)) {
						t.Errorf("Expected %v, got %v", tt.expected, result)
					}
				} else if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
				}
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}

func TestValidateFields(t *testing.T) {
	tests := []struct {
		name        string
		data        map[string]any
		fields      []logformat.LogFormatField
		expectError bool
	}{
		{
			name: "All required fields present",
			data: map[string]any{
				"timestamp": "2025-08-06",
				"status":    200,
			},
			fields: []logformat.LogFormatField{
				{RawName: "timestamp", SemanticName: strPtr("timestamp"), Required: true},
				{RawName: "status", SemanticName: strPtr("status"), Required: true},
			},
			expectError: false,
		},
		{
			name: "Missing required field",
			data: map[string]any{
				"status": 200,
			},
			fields: []logformat.LogFormatField{
				{RawName: "timestamp", SemanticName: strPtr("timestamp"), Required: true},
				{RawName: "status", SemanticName: strPtr("status"), Required: true},
			},
			expectError: true,
		},
		{
			name: "Required field is null",
			data: map[string]any{
				"timestamp": nil,
			},
			fields: []logformat.LogFormatField{
				{RawName: "timestamp", SemanticName: strPtr("timestamp"), Required: true},
			},
			expectError: true,
		},
		{
			name: "Optional field missing",
			data: map[string]any{
				"timestamp": "2025-08-06",
			},
			fields: []logformat.LogFormatField{
				{RawName: "timestamp", SemanticName: strPtr("timestamp"), Required: true},
				{RawName: "level", SemanticName: strPtr("level"), Required: false},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFields(tt.data, tt.fields)
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
