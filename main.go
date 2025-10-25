package main

import (
	"context"
	"gin-scaffold/cmd"
	"gin-scaffold/pkg/logging"
	"github.com/urfave/cli/v2"
	"os"
)

// VERSION You can specify the version number by compiling：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "v0.0.1"

//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@SecurityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	app := cli.NewApp()
	app.Name = "gin-scaffold"
	app.Version = VERSION
	ctx := logging.NewTagContext(context.Background(), "_main_")
	app.Commands = []*cli.Command{
		cmd.Start(ctx, VERSION),
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
