package utils

import (
	"encoding/json"
)

func IsValidJSON(input []byte) bool {
	var jw json.RawMessage
	return json.Unmarshal(input, &jw) == nil

}
