package utils

import (
	"os"
	"strings"
)

func ReadInput(file string) []string {
	raw_input, err := os.ReadFile(file)
	Check(err)

	lines := strings.Split(string(raw_input), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
