package main

import (
	"log"
	"flag"

	"github.com/rjkroege/id3dumper/tags"
)


var debug = flag.Bool("debug", false, "Set to true for more verbose debugging")


func main() {
	log.Println("hello")

	// Option parsing
	flag.Parse()
	
	for _, f := range flag.Args() {
		log.Println("handling", f)
		mdrd := tags.Match(f, *debug)
		if mdrd != nil {
			tag, err := mdrd.Get(f)
			if err != nil {
				log.Println("Skipping unreadable tag:", err)
				continue
			}
			log.Println(tag.String())
		}
	}
}


