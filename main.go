package main

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/cmd"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/urfave/cli/v2"
	"os"
)

// VERSION You can specify the version number by compiling：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "v0.0.1"

func main() {

	app := cli.NewApp()
	app.Name = "react-gin-ddd-admin"
	app.Version = VERSION
	app.Usage = ""

	ctx := logging.NewTagContext(context.Background(), "_main_")

	app.Commands = []*cli.Command{
		cmd.Start(ctx, VERSION),
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
