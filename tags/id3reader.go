package tags

import (
	"fmt"
	"log"
	"strconv"

	// The v2 suffix is for version2 of the module system?
	"github.com/bogem/id3v2/v2"
)

type id3 struct {
	debug bool
}

// There is a challenge here that I've not addressed: I've discovered the
// tags that I will process empirically. There will be some more to
// handle. There will be tracks with invalid tags. I need to discover
// this.

// Maybe dump the

func (tr *id3) Get(path string) (*Info, error) {

	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		return nil, fmt.Errorf("can't opening mp3 file %q: %v", path, err)
	}
	defer tag.Close()

	if tr.debug {
		frames := tag.AllFrames()
		for k, v := range frames {
			// TODO(rjk): should print a picture using the terminal capabilities
			switch {
			case k == "APIC":
				log.Println("APIC (i.e. image) present, image data elided")
			case k == "CHAP":
				log.Println("CHAP (i.e. chapter) present, chapter data elided")
			default:
				log.Printf("%s: %s\n", k, v)
			}
		}
	}

	// need to populate the Infos.
	// need to read the library... sigh.

	// recognize the year
	year, err := strconv.Atoi(tag.Year())
	if err != nil {
		return nil, fmt.Errorf("can't parse year %q from %q: %v", tag.Year(), path, err)
	}
	// recognize the tracknum
	track := tag.GetTextFrame("TRCK").Text
	ntrack, err := strconv.Atoi(track)
	if err != nil {
		return nil, fmt.Errorf("can't parse track %q from %q: %v", ntrack, path, err)
	}

	// What about fixing up the metadata?
	// What about attempting to set the metadata?

	// TODO(rjk): Extract the disk index (TPOS?)
	// TODO(rjk): Handle pictures.
	// TODO(rjk): Genre? Check that? Validate?

	return &Info{
		Author:    tag.Artist(),
		BookTitle: tag.Album(),
		//		SeriesTitle: // TIT2
		//		SeriesIndex: // ?
		Filename:   path,
		Year:       year,
		TrackIndex: ntrack,
		DiskIndex:  1,
		TrackName:  tag.Title(),
	}, nil
}

// TODO(rjk): Better display of pictures
// TODO(rjk): Better display of Chapter info
