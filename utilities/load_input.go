package utilities

import (
	"os"
	"strings"
)

// load input.txt, return a slice of strings with each line as an element
func LoadInput(filename string) []string {
	file, _ := os.ReadFile(filename)
	return strings.Split(string(file), "\n")
}
