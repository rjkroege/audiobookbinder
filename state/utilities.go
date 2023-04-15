package state

import (
	"strconv"
	"strings"
)

// Pretty print an Track.
// Would this be better in state? Probably.
func (i *Track) String() string {
	b := new(strings.Builder)

	b.WriteString("{\n")
	b.WriteString("	Author:		")
	b.WriteString(i.Author)
	b.WriteByte('\n')
	b.WriteString("	Booktitle:	")
	b.WriteString(i.Booktitle)
	b.WriteByte('\n')
	b.WriteString("	Filename:		")
	b.WriteString(i.Filename)
	b.WriteByte('\n')

	b.WriteString("	Year:		")
	b.WriteString(strconv.FormatInt(int64(i.Year), 10))
	b.WriteByte('\n')

	b.WriteString("	Track:		")
	b.WriteString(strconv.FormatInt(int64(i.Trackindex), 10))
	b.WriteByte('\n')

	b.WriteString("	Trackname:	")
	b.WriteString(i.Trackname)
	b.WriteByte('\n')

	b.WriteString("}")

	return b.String()
}
