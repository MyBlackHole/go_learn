package emulator

import (
    "os"
)

type serverCtxt struct {
	Port  int
	Dir   string
	Debug bool
}

var (
	globalServerCtxt serverCtxt

	globalOSSignalCh        = make(chan os.Signal, 1)
)
