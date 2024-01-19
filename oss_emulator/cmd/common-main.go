package emulator

import (
	cli "github.com/urfave/cli/v2"
    "context"
)

func buildServerCtxt(ctx *cli.Context, ctxt *serverCtxt) (err error) {
	ctxt.Port = ctx.Int("port")
	ctxt.Disk = ctx.String("disk")
	return
}

func contextCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
