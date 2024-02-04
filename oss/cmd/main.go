package oss

import (
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

func newApp(name string) *cli.App {
	commands := []*cli.Command{}
	commands = append(commands, serverCmd)

	app := cli.NewApp()
	app.Name = name
	app.HideHelpCommand = true
	app.Commands = commands
    app.OnUsageError = func(cCtx *cli.Context, err error, isSubcommand bool) error {
        panic(err)
    }

	return app
}

func Main(args []string) {
	appName := filepath.Base(args[0])

	if err := newApp(appName).Run(args); err != nil {
		os.Exit(1)
	}
}
