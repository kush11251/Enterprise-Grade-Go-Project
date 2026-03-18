package utils

import (
	"encoding/json"
	"fmt"
)

// ToJSON marshals the input data to JSON.
func ToJSON(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}