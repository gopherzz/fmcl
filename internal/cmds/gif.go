package cmds

import (
	"os/exec"

	"github.com/urfave/cli/v2"
)

var Gif *cli.Command = &cli.Command{
	Name:        "gif",
	Aliases:     []string{""},
	Usage:       "gif filename1 filename2",
	Description: "converts image to gif",
	Action:      action,
}

func action(ctx *cli.Context) error {
	exec.Command("ffmpeg", "-loglevel", "quiet", "-i", ctx.Args().First(), ctx.Args().Get(1)+".gif")
	return nil
}
