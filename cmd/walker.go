package cmd

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/rjkroege/audiobookbinder/global"
	"github.com/rjkroege/audiobookbinder/state"
	"github.com/rjkroege/audiobookbinder/tags"
)

func perfd(gctx *global.Context, queries *state.Queries, path string, info fs.FileInfo, err error) error {
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

		// I'd have to create a different context to do this in a go routine?
		if err := queries.CreateTrack(context.Background(), state.CreateTrackParams{
			Author:     tag.Author,
			Booktitle:  tag.Booktitle,
			Trackindex: tag.Trackindex,
			Year:       tag.Year,
			Diskindex:  tag.Diskindex,
			Filename:   tag.Filename,
			Trackname:  tag.Trackname,
		}); err != nil {
			log.Printf("can't insert %v into db because: %v", tag, err)
		}
	}
	return nil
}

// WalkAll walks the provided directory path looking for audiobook files.
// TODO(rjk): Add these to a database.
// TODO(rjk): put the loop into this.
func WalkAll(gctx *global.Context, root string) error {
	if err := state.OpenDb(gctx); err != nil {
		return fmt.Errorf("WalkAll can't open database %q: %v", gctx.Dbname, err)
	}
	queries := state.New(gctx.Db)

	// TODO(rjk): Should this be the io/fs.WalkDir to work on Windows?
	// TODO(rjk): Should this be a WalkDir invocation? Maybe that's faster?
	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		return perfd(gctx, queries, path, info, err)
	}); err != nil {
		return fmt.Errorf("can't walk %q: %v", root, err)
	}
	return nil
}

// 3 basic commands
// walk, report, join
// 2 advanced commands: fixtag, rip?
// how to implement tag fixing? TUI? A spreadsheet? (Go journal about that)
