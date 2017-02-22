package attachment

import (
	"encoding/json"
	"reflect"
	"strings"
)

// Marshal marshals SenderInfo
func (senderInfo *SenderInfo) Marshal() ([]byte, error) {
	return json.Marshal(senderInfo)
}

// Map transforms SenderInfo to map[string]string for embedding in
// Transaction/Operation struct.
func (senderInfo SenderInfo) Map() map[string]string {
	m := make(map[string]string)
	st := reflect.TypeOf(senderInfo)
	v := reflect.ValueOf(senderInfo)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			tagElements := strings.Split(jsonTag, ",")
			value := reflect.Indirect(v).FieldByName(field.Name).String()
			if value != "" {
				m[tagElements[0]] = value
			}
		}
	}
	return m
}
