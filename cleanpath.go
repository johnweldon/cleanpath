package cleanpath

import (
	"path"
	"strings"
)

// Clean accepts an array of strings representing PATH elements
// and returns a string with the elements joined by the os specific
// path separator, and with the duplicates removed, without losing
// initial order.
func Clean(sep rune, paths ...string) string {
	pos := 0
	working := make(map[string]int)
	for _, pth := range paths {
		for _, segment := range strings.FieldsFunc(pth, func(c rune) bool { return c == sep }) {
			clean := path.Clean(segment)
			// ignore relative paths
			if !path.IsAbs(clean) {
				continue
			}
			// Make sure relative order remains, and if duplicates then keep the first position
			if _, present := working[clean]; !present {
				working[clean] = pos
				pos++
			}
		}
	}
	array := make([]string, len(working))
	for segment, position := range working {
		array[position] = segment
	}
	return strings.Join(array, string(sep))
}
