package commands

import (
	"github.com/alexandervantrijffel/aws-appstream-client/pkg/awsclient"
	"github.com/alexandervantrijffel/goutil/logging"
	"github.com/urfave/cli/v2"
)

var StreamingURLCommand = &cli.Command{

	Name:      "streamingurl",
	Usage:     "create streaming url",
	ArgsUsage: "[streamingurl]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "stack-name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "fleet-name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user-id",
			Required: true,
		},
		&cli.StringFlag{
			Name: "application-id",
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

		url, err := awsclient.CreateStreamingURL(context, ctx.String("user-id"), ctx.String("application-id"), int64(60))
		if err != nil {
			return err
		}
		logging.Info("Streaming URL ", url)
		return nil

	},
}
