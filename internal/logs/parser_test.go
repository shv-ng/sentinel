package logs

import (
	"reflect"
	"testing"
)

func TestRegexParser(t *testing.T) {
	tests := []struct {
		name     string
		logline  string
		regex    string
		expected map[string]string
	}{
		{
			name:    "Parse timestamp and level",
			logline: "2025-08-06 INFO Starting service",
			regex:   `(?P<date>\d{4}-\d{2}-\d{2}) (?P<level>[A-Z]+)`,
			expected: map[string]string{
				"date":  "2025-08-06",
				"level": "INFO",
			},
		},
		{
			name:    "Parse IP and port",
			logline: "Connected to 192.168.1.1:8080",
			regex:   `(?P<ip>\d+\.\d+\.\d+\.\d+):(?P<port>\d+)`,
			expected: map[string]string{
				"ip":   "192.168.1.1",
				"port": "8080",
			},
		},
		{
			name:     "No match found",
			logline:  "invalid line",
			regex:    `(?P<key>\w+):(?P<value>\w+)`,
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RegexParser(tt.logline, tt.regex)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RegexParser(%q, %q) = %v; want %v", tt.logline, tt.regex, result, tt.expected)
			}
		})
	}
}

func TestJSONParser(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    map[string]any
		expectError bool
	}{
		{
			name:  "Valid JSON object",
			input: `{"name": "Alice", "age": 25}`,
			expected: map[string]any{
				"name": "Alice",
				"age":  float64(25),
			},
			expectError: false,
		},
		{
			name:        "Invalid JSON syntax",
			input:       `{"name": "Bob", "age":}`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty JSON object",
			input:       `{}`,
			expected:    map[string]any{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := JSONParser(tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got error: %v", tt.expectError, err)
				return
			}
			if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("JSONParser(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
