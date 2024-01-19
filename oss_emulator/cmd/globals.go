package emulator

import (
	"os"
)

type serverCtxt struct {
	Port  int
	Disk  string
	Debug bool
}

var (
	globalServerCtxt serverCtxt

	globalOSSignalCh = make(chan os.Signal, 1)

	globalLocalDrive StorageAPI
)
