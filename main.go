package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/rjkroege/audiobookbinder/cmd"
	"github.com/rjkroege/audiobookbinder/global"
	"github.com/rjkroege/audiobookbinder/tags"
)

var CLI struct {
	Debug bool `help:"Enable debugging conveniences as needed."`

	// TODO(rjk): Later, keep this somewhere smarter?
	Db   string `help:"Enable debugging conveniences as needed." default:"state.db"`
	Scan struct {
		Paths []string `arg:"" name:"path" help:"Paths to scan." type:"path"`
	} `cmd:"" help:"Scan directories for audiobook segments."`

	Report struct {
	} `cmd:"" help:"Print out a report about previously scanned audiobook segments."`

	Tagprint struct {
		Paths []string `arg:"" name:"path" help:"Paths to scan." type:"path"`
	} `cmd:"" help:"Print out a detailed list of tags in a particular track."`

	Reset struct {
	} `cmd:"" help:"Create a brand new empty database."`
}

func main() {
	log.Println("hello 1")
	ctx := kong.Parse(&CLI)

	cmdctx := &global.Context{
		Debug:  CLI.Debug,
		Dbname: CLI.Db,
	}

	switch ctx.Command() {
	case "scan <path>":
		for _, f := range CLI.Scan.Paths {
			err := cmd.WalkAll(cmdctx, f)
			if err != nil {
				log.Println("Walk failed", err)
			}
		}
	case "report":
		log.Fatal("report functionality is not yet implemented")
	case "tagprint <path>":
		md, err := tags.MakeMetaReader(cmdctx.Debug)
		if err != nil {
			log.Fatalf("tagprint can't make a MetaReader because: %v", err)
		}
		defer md.Clunk()
		for _, f := range CLI.Tagprint.Paths {
			if err := md.Tagprint(f); err != nil {
				log.Printf("can't read %q: %v\n", f, err)
			}
		}

	case "reset":
		if err := cmd.ResetDatabase(cmdctx); err != nil {
			log.Fatalf("Can't reset database %q: %v", cmdctx.Dbname, err)
		}
	default:
		log.Fatal("Missing command: ", ctx.Command())
	}
}
