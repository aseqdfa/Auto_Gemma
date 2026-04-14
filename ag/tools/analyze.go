package tools

import (
	"os"
)

func Analyze(path string) string {

	files, _ := os.ReadDir(path)

	out := "Files:\n"
	for _, f := range files {
		out += f.Name() + "\n"
	}

	return out
}
