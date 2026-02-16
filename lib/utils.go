package lib

import (
	"strings"
)

func getPaylaod(arr []string) string {
	var payload string
	command := strings.Join(arr, " ")
	start := strings.IndexAny(command, "{")
	if start != -1 {
		end := strings.IndexAny(command, "}")
		payload = command[start+1 : end]
	}

	return payload
}
