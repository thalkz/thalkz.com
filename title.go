package main

import (
	"strings"
)

func toFilename(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func capitalizeWord(word string) string {
	capitalizedToken := []byte(word)
	capitalizedToken[0] = word[0] + 'A' - 'a'
	return string(capitalizedToken)
}

func capitalizeAll(words []string) []string {
	capitalized := make([]string, len(words))
	for i := range words {
		capitalized[i] = capitalizeWord(words[i])
	}
	return capitalized
}

func toTitle(filepath string) string {
	filename := toFilename(filepath)
	words := strings.Split(filename, "_")
	capitalized := capitalizeAll(words)
	return strings.Join(capitalized, " ")
}
