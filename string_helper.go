package main

// In Go, if we have multiple files with the same package, they are all visible
// If there are different packages, we make things visible across them capitalizing
// the first letter of functions or elements on structs

import (
	"fmt"
	"strings"
)

// Variatic Variable as parameter here
func containsEmpty(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

func sanitizeLine(lineRow string) string {
	lineRow = strings.ReplaceAll(lineRow, "\"", "'")
	lineRow = strings.ReplaceAll(lineRow, ",", ".")

	return lineRow
}

func emptyLine() {
	fmt.Printf("\n")
}
