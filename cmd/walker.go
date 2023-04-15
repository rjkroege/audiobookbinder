package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/rjkroege/id3dumper/tags"
		"github.com/rjkroege/id3dumper/global"
)

func perfd(gctx *global.Context, path string, info fs.FileInfo, err error) error {
	// Directories can't ever be of interest.
	if info.IsDir() {
		return nil
	}

	log.Println("handling", path)
	mdrd := tags.Match(path, gctx.Debug)
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

// WalkAll walks the provided directory path looking for audiobook files.
// TODO(rjk): Add these to a database.
// TODO(rjk): put the loop into this.
func WalkAll(gctx *global.Context, root string) error {
	// TODO(rjk): Should this be the io/fs.WalkDir to work on Windows?
	// TODO(rjk): Should this be a WalkDir invocation? Maybe that's faster?
	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		return perfd(gctx, path, info, err)
	}); err != nil {
		return fmt.Errorf("can't walk %q: %v", root, err)
	}
	return nil
}

// 3 basic commands
// walk, report, join
// 2 advanced commands: fixtag, rip?
// how to implement tag fixing? TUI? A spreadsheet? (Go journal about that)
