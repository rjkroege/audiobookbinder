package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/rjkroege/audiobookbinder/cmd"
	"github.com/rjkroege/audiobookbinder/global"
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

	Reset struct {
	} `cmd:"" help:"Create a brand new empty database."`
}

func main() {
	log.Println("hello 1")

	ctx := kong.Parse(&CLI)

	log.Println("the db file", CLI.Db)

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
	case "reset":
		if err := cmd.ResetDatabase(cmdctx); err != nil {
			log.Fatalf("Can't reset database %q: %v", cmdctx.Dbname, err)
		}
	default:
		log.Fatal("Missing command: ", ctx.Command())
	}
}
