package tags

import (
	"path/filepath"
	"strings"
	"strconv"
)


// Info is the cannonical structure for metadata.
// TODO(rjk): 
type Info struct {
	// Author of this book
	Author string

	// Title of this book
	BookTitle string

	// Title of the series (if part of a series)
	SeriesTitle string

	// Position of fragment within book
	DiskIndex int
	TrackIndex int

	// Year of publication
	Year int

	// Where it is on the disk
	Filename string

	// Track name
	TrackName string

	// Metadata is valid
	CompleteMetadata bool

	// TODO(rjk): picture? how to deal
	// chapter data (how to deal)
}

// Reader providees a standard way for per-file tag reading mechanisms.
type MetaReader interface {
// Reads the tag data from file name and returns it in the cannonical
// Info structure or returns an error if impossible to do so.
// TODO(rjk): API second-guessing: it's possible that I could save the
// path into the object? Shrug. It doesn't matter at this level of complexity.
	Get(path string) (*Info, error)

}

// Pretty print an Info.
func (i *Info) String() string {
	b := new(strings.Builder)
	
	b.WriteString("{\n")
	b.WriteString("	Author:		")
	b.WriteString(i.Author)
	b.WriteByte('\n')
	b.WriteString("	BookTitle:	")
	b.WriteString(i.BookTitle)
	b.WriteByte('\n')
	b.WriteString("	Filename:		")
	b.WriteString( i.Filename)
	b.WriteByte('\n')

	b.WriteString("	Year:		")
	b.WriteString( strconv.FormatInt( int64( i.Year ), 10))
	b.WriteByte('\n')

	b.WriteString("	Track:		")
	b.WriteString( strconv.FormatInt( int64( i.TrackIndex ), 10))
	b.WriteByte('\n')

	b.WriteString("	TrackName:	")
	b.WriteString( i.TrackName )
	b.WriteByte('\n')

	b.WriteString("}")

	return b.String()
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
