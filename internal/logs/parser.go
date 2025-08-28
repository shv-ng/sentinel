package logs

import (
	"encoding/json"
	"regexp"
)

func RegexParser(logline, regex string) map[string]string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(logline)
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" && i < len(match) {
			result[name] = match[i]
		}
	}
	return result
}

func JSONParser(jsoninput string) (map[string]any, error) {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(jsoninput), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
