package cmd

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/internal/bootstrap"
	"github.com/urfave/cli/v2"
)

func Start(ctx context.Context, version string) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Required: true,
				Aliases:  []string{"c"},
				Usage:    "Runtime configuration",
			},
			&cli.StringFlag{
				Name:    "static",
				Aliases: []string{"s"},
				Usage:   "Static files directory",
			},
		},
		Action: func(c *cli.Context) error {
			return bootstrap.RunServer(ctx,
				&bootstrap.Options{
					ConfigFileDir: c.String("conf"),
					Version:       version,
					StaticFileDir: c.String("static"),
				})

		},
	}
}
