package utils

import (
	"encoding/json"
)

func JsonFormater(jsonByte []byte) map[string]interface{} {
	var data map[string]interface{}

	json.Unmarshal([]byte(jsonByte), &data)

	return data
}