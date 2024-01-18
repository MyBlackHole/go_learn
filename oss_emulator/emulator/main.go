package emulator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	cli "github.com/urfave/cli/v2"
)


var ServerFlags = []cli.Flag{
	&cli.StringFlag{
		Name:   "port, p",
		Value:  DefaultPort,
		Usage:  "监听端口设置",
	},
	&cli.StringFlag{
		Name:   "dir, d",
		Value:  DefaultDir,
		Usage:  "挂载目录",
	},
	&cli.BoolFlag{
		Name:   "debug",
		Usage:  "调试",
		Hidden: true,
	},
}


var HelpTemplate = `NAME:
  {{.Name}} - {{.Usage}}

DESCRIPTION:
  {{.Description}}

USAGE:
  {{.HelpName}} {{if .VisibleFlags}}[FLAGS] {{end}}COMMAND{{if .VisibleFlags}}{{end}} [ARGS...]

COMMANDS:
  {{range .VisibleCommands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
  {{end}}{{if .VisibleFlags}}
FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}
VERSION:
  {{.Version}}
`


func newApp(name string) *cli.App {
	commands := []*cli.Command{}
    commands = append(commands, serverCmd)

	cli.VersionPrinter = printVersion

	app := cli.NewApp()
	app.Name = name
	app.Version = ReleaseTag
	app.Usage = "High Performance Object Storage"
	app.Description = `Build high performance data infrastructure for machine learning, analytics and application data workloads with MinIO`
	app.Flags = ServerFlags
	app.HideHelpCommand = true // Hide `help, h` command, we already have `minio --help`.
	app.Commands = commands
	app.CustomAppHelpTemplate = HelpTemplate

	return app
}

func versionBanner(c *cli.Context) io.Reader {
	banner := &strings.Builder{}
	fmt.Fprintf(banner, "%s version %s (commit-id=%s)\n", c.App.Name, c.App.Version, CommitID)
	return strings.NewReader(banner.String())
}

func printVersion(c *cli.Context) {
	_,_ = io.Copy(c.App.Writer, versionBanner(c))
}

func Main(args []string) {
	appName := filepath.Base(args[0])

	if err := newApp(appName).Run(args); err != nil {
		os.Exit(1)
	}
}
