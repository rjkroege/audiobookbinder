# TL;DR
To become a tool that will join audio books divided up into tracks into single large
audio files.

# Rationale
It's hard to know how to manage audiobooks in iOS with the
introduction of Apple Music. Where do they fit? I copy an audiobook CD
into `Music` and then play it `Music`? Or `Books` with the purchased
audiobooks? Moreover, there is no "matched" functionality for
audiobooks -- they all end up in the "uploaded" state and this has its
problems (that I attemped to fix with [rjkroege/kickmusicmatch: Force
Apple Music to try to re-upload previously erroring
tracks](https://github.com/rjkroege/kickmusicmatch))

So I had the idea that I could listen to them in VLC app. And large
puddles of tracks seemed tedious so I thought that single large
whole-book audio files would be easier to manage. And then I started
writing a tool to join the audiobook tracks together.

# Discussion
First, it's conceivable that this entire undertaking was a very poor use of my time for the following
reasons:

* VLC can play groups of tracks
* AirDrop can copy a group of tracks into VLC
* Limited experience with Apple Music matching in Ventura suggests
that Apple has made the matching feature less flaky than the bugs that
prompted to write `kickmusicmatch` so simply upgrading would (at
least) ameliorate the issue.

Ignoring that all of this work was perhaps completely pointless, I did
learn a number of useful things here and this README serves to record
them before I forget.

## Lesson: binding
I found this cool tool [crra/mp3binder: 🎵 concatenating, joining,
binding MP3 files without
re-encoding](https://github.com/crra/mp3binder) that can join mp3
files together. I manually concatenanted some audiobooks and tried
them out as single files in VLC. That worked fine.

I (perhaps foolishly) ignored what I'd do with `.m4a` etc. files
trusting that I could make my tool have some kind of nice Go interface
for "join these tracks" and I'd end up needing to write something
different for each kind of file.

## Lesson: ID3 Parsing
But `mp3binder` needs (as one might expect) to know which files to
bind and the metadata to set on the result. So to successfully drive
`mp3binder`, `audiobookbinder` needs to read the metadata from the
constiutent music files, group and order the files by album and
invokve `mp3binder`. So I had to parse audio files for their metadata.

* It was an unpleasant discovery how complicated ID3 parsing can be.
* I attempted to use several different Go ID3 libraries. None operated
as I understood their respective READMEs.
* If the library supports reading and writing ID3 tags, *test in advance* that the library can successfully
read its own output.
* In particular: 
* As with binding, I (foolishly) blithely ignored file types other than MP3.

There is a general lesson to be had had here that I should try harder
to follow: write tests for third party code before committing to an
implementation that depends on it. Doh.

## Lesson: Sqlite State Management
I didn't want to produce incomplete single-track audio books. I wanted
to track state across runs of the tools: which tracks made a complete
book, which ones needed metadata correction, which ones had been
emitted. I decided that I could express the idea of complete books
nicely with SQL queries. I hence needed to learn about the Go SQL
scene. Here my surprises were all pleasant:

* [GitHub - sqlc-dev/sqlc: Generate type-safe code from SQL](https://github.com/sqlc-dev/sqlc) is
a very cool tool for connecting SQL databases to Go code: just write some SQL and use the
typesafe Go bindings.
* [sqlite package - modernc.org/sqlite - Go Packages](https://pkg.go.dev/modernc.org/sqlite?utm_source=godoc) really is a drop-in pure Go replacement for Sqlite3. And it works really well.

# Summary

* don't build something that you don't need to build at all
* test all of your dependencies before starting
