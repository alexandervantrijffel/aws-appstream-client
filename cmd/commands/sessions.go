package commands

import (
	"github.com/alexandervantrijffel/aws-appstream-client/pkg/awsclient"
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/urfave/cli/v2"
)

var SessionsCommand = &cli.Command{

	Name:      "sessions",
	Usage:     "describe-sessions",
	ArgsUsage: "[sessions]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "stack-name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "fleet-name",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		service, err := awsclient.NewService()
		if err != nil {
			return err
		}
		context := awsclient.AppStreamContext{
			Service:   service,
			StackName: ctx.String("stack-name"),
			FleetName: ctx.String("fleet-name"),
		}

		activeSessions, err := awsclient.DescribeSessions(context)
		if err != nil {
			return err
		}

		if len(activeSessions) == 0 {

			logging.Info("No active AppStream sessions found for the current region")
		} else {
			logging.Info("Sessions", activeSessions)
		}
		return nil
	},
}
