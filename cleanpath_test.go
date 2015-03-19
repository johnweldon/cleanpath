package cleanpath_test

import (
	"os"
	"testing"

	"gopkg.in/johnweldon/cleanpath.v0"
)

func TestCleanpath(t *testing.T) {
	for input, expected := range testCase1 {
		got := cleanpath.Clean(os.PathListSeparator, []string{input}...)
		if expected != got {
			t.Error("Expected '" + expected + "', got '" + got + "'\n")
		}
	}
}
