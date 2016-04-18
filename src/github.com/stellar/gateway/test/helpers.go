package test

import (
	"encoding/json"
)

// StringToJSONMap transforms
func StringToJSONMap(value string) (m map[string]interface{}) {
	json.Unmarshal([]byte(value), &m)
	return
}
