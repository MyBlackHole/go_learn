package emulator

import (
	"os"

	"github.com/nutsdb/nutsdb"
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

	globalMetaDb *nutsdb.DB

	globalOssDefaultOwnerID = "00220240124"

	globalDirSuffix = "__XLDIR__"
)
