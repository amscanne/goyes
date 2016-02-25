// Binary goyes is a Go port of yes.
package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
	"time"
)

var delay = flag.Duration("delay", 0, "delay between lines")

func main() {
	// Parse flags.
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] [STRING]...\n", os.Args[0])
		flag.PrintDefaults()
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
		if *delay != 0 {
			time.Sleep(*delay)
		}
	}
}
