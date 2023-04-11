package main

import (
	"flag"
	"log"

	"github.com/rjkroege/id3dumper/cmd"
)

// TODO(rjk): Consider switching to kong?
func main() {
	log.Println("hello")

	// Option parsing.
	flag.Parse()

	for _, f := range flag.Args() {
		err := cmd.WalkAll(f)
		if err != nil {
			log.Println("Walk failed", err)
		}
	}
}
