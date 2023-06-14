package fmcl

import (
	"github.com/gopherzz/fmcl/internal/cmds"
	"github.com/urfave/cli/v2"
)

type fmcl struct {
	app *cli.App
}

func NewFmcl(app *cli.App) *fmcl {
	return &fmcl{
		app: app,
	}
}

func (c *fmcl) Start(args []string) error {
	c.initCommands()
	return c.app.Run(args)
}

func (c *fmcl) initCommands() {
	c.app.Commands = []*cli.Command{
		cmds.Gif,
	}
}
