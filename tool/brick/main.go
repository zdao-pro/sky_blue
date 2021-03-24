package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Println(Version)
	app := cli.NewApp()
	app.Name = "brick"
	app.Usage = "brick工具集"
	app.Version = Version
	app.Commands = []*cli.Command{
		{
			Name:            "new",
			Aliases:         []string{"n"},
			Usage:           "创建新项目",
			Action:          runNew,
			SkipFlagParsing: true,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "brick build",
			Action:  buildAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "brick run",
			Action:  runAction,
		},
		{
			Name:            "tool",
			Aliases:         []string{"t"},
			Usage:           "brick tool",
			Action:          toolAction,
			SkipFlagParsing: true,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "brick version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
		{
			Name:   "self-upgrade",
			Usage:  "brick self-upgrade",
			Action: upgradeAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func runNew(ctx *cli.Context) error {
	return installAndRun("genproject", ctx.Args().Slice())
}
