package main

import (
	"os"

	"github.com/alexandervantrijffel/aws-appstream-client/cmd/commands"
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "aws-appstream-client",
		HelpName:             "aws-appstream-client",
		EnableBashCompletion: true,
		Usage:                "AWS AppStream 2.0 cli",
		Commands: []*cli.Command{
			commands.SessionsCommand,
			commands.StreamingURLCommand,
		},
	}

	logging.InitWith(app.Name, false)

	if err := app.Run(os.Args); err != nil {
		logging.Fatal(err)
	}
}
