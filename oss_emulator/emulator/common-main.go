package emulator

import (
	cli "github.com/urfave/cli/v2"
)

func buildServerCtxt(ctx *cli.Context, ctxt *serverCtxt) (err error) {
	ctxt.Port = ctx.Int("port")
	return
}
