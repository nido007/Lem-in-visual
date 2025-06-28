package main

import (
	"strings"
)

// parseInput splits the raw input text into individual lines, trimming any leading or trailing whitespace.
// Returns a slice of strings, each representing one line of the original input.
func parseInput(input string) []string {
	// Remove any extra spaces at the beginning and end
	trimmed := strings.TrimSpace(input)

	// Split the text into separate lines
	lines := strings.Split(trimmed, "\n")

	// Clean up each line by removing extra spaces
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}
