package tags

import (
	//	"path/filepath"
	"fmt"

	"github.com/barasher/go-exiftool"
	"github.com/rjkroege/audiobookbinder/state"
)

// Reader providees a standard way for per-file tag reading mechanisms.
type MetaReader interface {
	// Reads the tag data from file name and returns it in the cannonical
	// Info structure or returns an error if impossible to do so.
	// TODO(rjk): API second-guessing: it's possible that I could save the
	// path into the object? Shrug. It doesn't matter at this level of complexity.
	Get(path string) (*state.Track, error)

	// Prints the detailed list of tag contents.
	Tagprint(path string) error

	// Discards a previously built MetaReader
	Clunk() error
}

// Creates a metadata reader based on ExifTool. ExifTool can parse all
// the tags so build one of these.
func MakeMetaReader(debug bool) (MetaReader, error) {

	et, err := exiftool.NewExiftool()
	if err != nil {
		return nil, fmt.Errorf("Can't make a MetaReader because %v", err)
	}
	return &id3{
		debug: debug,
		exol:  et,
	}, nil
}
