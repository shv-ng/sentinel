package utils

import "testing"

func TestIsValidJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "Valid JSON object",
			input:    []byte(`{"name":"John","age":30}`),
			expected: true,
		},
		{
			name:     "Invalid JSON (missing brace)",
			input:    []byte(`{"name":"John"`),
			expected: false,
		},
		{
			name:     "Valid JSON array",
			input:    []byte(`[1, 2, 3]`),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidJSON(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidJSON(%s) = %v; want %v", string(tt.input), result, tt.expected)
			}
		})
	}
}
