package tags

import (
	"path/filepath"

	"github.com/rjkroege/id3dumper/state"	
)


// Reader providees a standard way for per-file tag reading mechanisms.
type MetaReader interface {
	// Reads the tag data from file name and returns it in the cannonical
	// Info structure or returns an error if impossible to do so.
	// TODO(rjk): API second-guessing: it's possible that I could save the
	// path into the object? Shrug. It doesn't matter at this level of complexity.
	Get(path string) (*state.Track, error)
}


// Return a MetaReader implementation appropriate to path's extension or nil.
// TODO(rjk): I note the possibility of supporting complex options with
// the Pike options pattern as needed.
func Match(path string, debug bool) MetaReader {
	switch filepath.Ext(path) {
	case ".mp3", ".MP3":
		return &id3{debug: debug}
	}
	return nil
}
