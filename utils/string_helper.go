package utils

// In Go, if we have multiple files with the same package, they are all visible
// If there are different packages, we make things visible across them capitalizing
// the first letter of functions or elements on structs

import (
	"fmt"
	"strings"
)

// ContainsEmpty - fails if string is empty
func ContainsEmpty(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

// SanitizeLine - Remove invalid characters to build a valid CSV row
func SanitizeLine(lineRow string) string {
	lineRow = strings.ReplaceAll(lineRow, "\"", "'")
	lineRow = strings.ReplaceAll(lineRow, ",", ".")

	return lineRow
}

// EmptyLine - Empty Line
func EmptyLine() {
	fmt.Printf("\n")
}
