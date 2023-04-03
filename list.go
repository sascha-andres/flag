package flag

import "strings"

// ListFromString returns a slice of strings from a comma-separated string.
func ListFromString(s string) []string {
	return strings.Split(s, ",")
}
