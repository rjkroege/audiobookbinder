package main

import (
	"log"
	"os"

	// The v2 suffix is for version2 of the module system?
	"github.com/bogem/id3v2/v2"
)

func main() {
	log.Println("hello")

	for _, f := range os.Args[1:] {
		log.Println("handling", f)
		readTags(f)
	}
}

func readTags(fn string) {
	tag, err := id3v2.Open(fn, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("Error while opening mp3 file: ", err)
	}
	defer tag.Close()

	frames := tag.AllFrames()
	
	for k, v := range frames {
		// TODO(rjk): should print a picture using the terminal capabilities
		if k == "APIC" {
			log.Println("APIC (i.e. image) present, image data elided")
		} else {
			log.Printf("%s: %s\n", k, v)
		}
	}
}

// TODO(rjk): Better display of pictures
// TODO(rjk): Better display of Chapter info
