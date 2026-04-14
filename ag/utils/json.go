package utils

import "encoding/json"

func Parse(s string, v interface{}) {
	json.Unmarshal([]byte(s), v)
}
