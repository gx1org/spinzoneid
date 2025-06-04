package app

import (
	"math/rand"
	"strings"
	"time"
)

func generateID() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func parseOptions(input string) []string {
	lines := strings.Split(input, "\n")
	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func pickRandomOption(options []string) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return options[rand.Intn(len(options))]
}
