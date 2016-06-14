package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/johnweldon/cleanpath"
)

var sep = separator{val: os.PathListSeparator}

type separator struct {
	val rune
}

func (s *separator) String() string {
	if s == nil {
		return ""
	}
	return string(s.val)
}

func (s *separator) Set(v string) error {
	for ix, r := range v {
		if ix > 0 {
			return fmt.Errorf("cannot use %q as a separator; only one rune allowed", v)
		}
		s.val = r
	}
	return nil
}

func init() {
	flag.Var(&sep, "separator", "separator defaults to os specific separator")
}

func main() {
	flag.Parse()
	args := flag.Args()
	clean := cleanpath.Clean(sep.val, args...)
	fmt.Fprintf(os.Stdout, "%s", clean)
}
