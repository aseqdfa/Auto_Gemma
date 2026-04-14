package tools

import "os"

func ReadFile(path string) string {
	b, _ := os.ReadFile(path)
	return string(b)
}
