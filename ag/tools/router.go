package tools

import "strings"

func FastRoute(input string) (string, bool) {

	if strings.HasPrefix(input, "/run") {
		return RunCmd(strings.TrimPrefix(input, "/run ")), true
	}

	if strings.HasPrefix(input, "/read") {
		return ReadFile(strings.TrimPrefix(input, "/read ")), true
	}

	return "", false
}
