package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/johnweldon/cleanpath.v0"
)

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	clean := cleanpath.Clean(args)
	fmt.Fprintf(os.Stdout, "%s", clean)
}
