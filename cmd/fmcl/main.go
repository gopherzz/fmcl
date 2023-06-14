package main

import (
	"log"
	"os"

	"github.com/gopherzz/fmcl/internal/fmcl"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "fmcl",
		Description: "FMCL - Command-line utility for better ffmpeg user experience",
	}
	fmcl := fmcl.NewFmcl(app)
	if err := fmcl.Start(os.Args); err != nil {
		log.Fatal(err)
	}
}
