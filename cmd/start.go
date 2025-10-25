package cmd

import (
	"context"

	"gin-scaffold/internal/initialize"
	"github.com/urfave/cli/v2"
)

func Start(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start the server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Required: true,
				Aliases:  []string{"c"},
				Usage:    "Runtime configuration dir",
			},
		},
		Action: func(c *cli.Context) error {
			return initialize.RunServer(ctx,
				&initialize.Options{
					ConfigFileDir: c.String("conf"),
					Version:       version,
				})

		},
	}
}
