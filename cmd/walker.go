package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/rjkroege/id3dumper/tags"
)

var debug = flag.Bool("debug", false, "Set to true for more verbose debugging")

func perfd(path string, info fs.FileInfo, err error) error {
	// Directiroes can't ever be of interest.
	if info.IsDir() {
		return nil
	}

	log.Println("handling", path)
	mdrd := tags.Match(path, *debug)
	if mdrd != nil {
		tag, err := mdrd.Get(path)
		if err != nil {
			log.Println("Skipping unreadable tag:", err)
			return nil
		}

		// Action is here
		log.Println(tag.String())
	}
	return nil
}

func WalkAll(root string) error {
	// TODO(rjk): Should this be the io/fs.WalkDir to work on Windows?
	// TODO(rjk): Should this be a WalkDir invocation? Maybe that's faster?
	if err := filepath.Walk(root, perfd); err != nil {
		return fmt.Errorf("can't walk %q: %v", root, err)
	}
	return nil
}
