package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/rjkroege/id3dumper/cmd"
)

var CLI struct {
	Debug bool `help:"Enable debugging conveniences as needed."`
	Scan  struct {
		Paths []string `arg:"" name:"path" help:"Paths to scan." type:"path"`
	} `cmd:"" help:"Scan directories for audiobook segments."`

	Report struct {
	} `cmd:"" help:"Print out a report about previously scanned audiobook segments."`
}

func main() {
	log.Println("hello")

	ctx := kong.Parse(&CLI)
	cmdctx := &cmd.Context{
		Debug: CLI.Debug,
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
	default:
		log.Fatal("Missing command: ", ctx.Command())

	}
}
