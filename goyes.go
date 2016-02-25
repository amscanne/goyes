// Binary goyes is a Go port of yes.
package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Parse no flags.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [STRING]...\n", os.Args[0])
	}
	flag.Parse()

	// Construct our string.
	args := flag.Args()
	s := "y"
	if len(args) > 0 {
		s = strings.Join(args, " ")
	}

	for {
		// Print it over and over again.
		fmt.Fprintf(os.Stdout, "%s\n", s)
	}
}
