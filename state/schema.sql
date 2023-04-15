-- TODO(rjk): It's highly likely that additional refinements are needed here.
-- TODO(rjk): Support for pictures.
-- Content extracted from metadata in each track.
CREATE TABLE IF NOT EXISTS tracks (
	id   INTEGER PRIMARY KEY,

	-- Author of this book.
	author text NOT NULL,

	-- Title of this book.
	booktitle text NOT NULL,

	-- Disk of set.
	diskindex INTEGER NOT NULL,

	-- Track of tracks on a specific disk.
	trackindex INTEGER NOT NULL,

	-- Year of this book.
	year INTEGER NOT NULL,

	-- Filename of this track.
	filename text UNIQUE NOT NULL,

	-- Name of this track.
	trackname text NOT NULL
);

CREATE TABLE IF NOT EXISTS  books (
	id   INTEGER PRIMARY KEY,
	complete INTEGER  NOT NULL
);

-- Note the presence of the foreign key constraint.
-- Will this be adequate for my sorting needs? I guess I'll find out.
CREATE TABLE  IF NOT EXISTS bookmembership (
	bookid INTEGER  NOT NULL REFERENCES books(id),
	trackid INTEGER NOT NULL REFERENCES tracks(id)
);

