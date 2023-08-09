package tags

import (
	"fmt"
	//	"log"
	"strconv"
	"strings"

	"github.com/barasher/go-exiftool"
	"github.com/rjkroege/audiobookbinder/state"
)

type id3 struct {
	debug bool
	exol  *exiftool.Exiftool
}

func (md *id3) Clunk() error {
	// TODO(rjk): I could consider making a better error message.
	return md.exol.Close()
}

func printAllTags(info exiftool.FileMetadata) {
	for k, v := range info.Fields {
		fmt.Printf("[%v] %v\n", k, v)
	}
}

// TODO(rjk): These require some kind of testing.
// parseTrack parses the "track of total track" two-tuple.
func parseTrack(info exiftool.FileMetadata) (int, int, error) {
	tagname := "TrackNumber"

	numoftotal, err := info.GetString(tagname)

	// The absence of some kind of track identity is a fatal error.
	// TODO(rjk): I may need to relax this error case once I permit editing the tag data.
	if err != nil && err == exiftool.ErrKeyNotFound {
		// Try Track instead
		tagname = "Track"
		numoftotal, err = info.GetString(tagname)
	}

	// This case indicates some kind of file corruption or I/O errors or there's no track
	// identity at all.
	if err != nil {
		return -1, -1, fmt.Errorf("can't get %s from track %q: %v", tagname, info.File, err)
	}

	n, t := -1, -1
	num := numoftotal
	div := strings.Index(numoftotal, " of ")
	if div > 0 {
		num = numoftotal[0:div]
		tot := numoftotal[div+len(" of "):]

		t, err = strconv.Atoi(tot)
		if err != nil {
			return -1, -1, fmt.Errorf("invalid %s value %q from track %q: %v", tagname, numoftotal, info.File, err)
		}
	}

	n, err = strconv.Atoi(num)
	if err != nil {
		return -1, -1, fmt.Errorf("invalid %s value %q from track %q: %v", tagname, numoftotal, info.File, err)
	}

	return n, t, nil
}

// parseDisk parses the "disk of total disks" two-tuple.
func parseDisk(info exiftool.FileMetadata) (int, int, error) {
	tagname := "DiskNumber"
	numoftotal, err := info.GetString(tagname)

	// The absence of the disk-tuple doesn't necessarily indicate that this
	// track is loose. Audiobooks comprised of a single disk of
	// pre-compressed data would appear (by inspection) to have no disk data.
	// In this case, I return 0,0 as a valid tuple. The presence of the 0,0
	// disktuple will signify to the later processing steps that we don't
	// know.
	if err != nil && err == exiftool.ErrKeyNotFound {
		return 0, 0, nil
	}

	// This case indicates some kind of file corruption or I/O errors. I
	// should ignore tracks in this situation.
	if err != nil {
		return -1, -1, fmt.Errorf("can't get %s from track %q: %v", tagname, info.File, err)
	}

	div := strings.Index(numoftotal, " of ")
	if div < 0 {
		return -1, -1, fmt.Errorf("invalid tag %s from track %q: %v", tagname, info.File, err)
	}
	num := numoftotal[0:div]
	tot := numoftotal[div+len(" of ")]

	n, err := strconv.Atoi(string(num))
	if err != nil {
		return -1, -1, fmt.Errorf("invalid tag %s from track %q: %v", tagname, info.File, err)
	}
	t, err := strconv.Atoi(string(tot))
	if err != nil {
		return -1, -1, fmt.Errorf("invalid tag %s from track %q: %v", tagname, info.File, err)
	}

	return n, t, nil
}

// There is a challenge here that I've not addressed: I've discovered the
// tags that I will process empirically. There will be some more to
// handle. There will be tracks with invalid tags. I need to discover
// this.

func (tr *id3) Get(path string) (*state.Track, error) {
	infos := tr.exol.ExtractMetadata(path)
	if len(infos) != 1 {
		return nil, fmt.Errorf("unexpected number of tracks from ExractMetadata")
	}
	info := infos[0]

	// Dump the info.
	if tr.debug {
		printAllTags(info)
	}

	filetype, err := info.GetString("FileType")
	if err != nil {
		return nil, fmt.Errorf("can't determine filetype from %q: %v", path, err)
	}

	// TODO(rjk): Expand the selection of file types.
	if filetype != "MP3" && filetype != "M4A" {
		return nil, fmt.Errorf("%q is unsupported type", path)
	}

	// Absence of year should not be fatal.
	year, err := info.GetInt("Year")
	if err != nil {
		year = 0
	}

	author, err := info.GetString("Artist")
	if err != nil {
		return nil, fmt.Errorf("can't parse Artist from %q: %v", path, err)
	}

	trackname, err := info.GetString("Title")
	if err != nil {
		return nil, fmt.Errorf("can't parse Title from %q: %v", path, err)
	}

	booktitle, err := info.GetString("Album")
	if err != nil {
		return nil, fmt.Errorf("can't parse Album from %q: %v", path, err)
	}

	// TODO(rjk): Record the disk and track totals into the database.
	// TODO(rjk): Attempt to guess the track number based on heuristics in the filename.
	// Observation: fixing up the data should be a separate step from loading the database.
	// TODO(rjk): Be more forgiving in parsing the the data here.
	disk, _, err := parseDisk(info)
	if err != nil {
		return nil, fmt.Errorf("no disk info: %v", err)
	}

	track, _, err := parseTrack(info)
	if err != nil {
		return nil, fmt.Errorf("no track info: %v", err)
	}

	// TODO(rjk): Stash the genre.
	//	_, err := info.GetString("Genre")
	//	if err != nil {
	//		return nil, fmt.Errorf("can't parse Genre from %q: %v", path, err)
	//	}

	// TODO(rjk): Handle pictures.

	return &state.Track{
		Author:    author,
		Booktitle: booktitle,
		//		SeriesTitle: // TIT2
		//		SeriesIndex: // ?
		Filename:   path,
		Year:       int64(year),
		Trackindex: int64(track),
		Diskindex:  int64(disk),
		Trackname:  trackname,
	}, nil

}

// TODO(rjk): Better display of pictures
// TODO(rjk): Better display of Chapter info

func (tr *id3) Tagprint(path string) error {
	// TODO(rjk): I am using this in a sync way. So it will be slower than
	// it could perhaps be. But I suspect that this won't matter very much.
	infos := tr.exol.ExtractMetadata(path)
	for _, fi := range infos {
		if fi.Err != nil {
			fmt.Printf("can't read %q: %v\n", fi.File, fi.Err)
			continue
		}

		for k, v := range fi.Fields {
			fmt.Printf("[%v] %v\n", k, v)
		}

	}

	return nil
}
