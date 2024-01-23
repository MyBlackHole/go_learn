package emulator

import (
	"fmt"
	"os"
)

func handleSignals() {
	exit := func(success bool) {
		if success {
			os.Exit(0)
		}

		os.Exit(1)
	}

	for {
		select {
		case osSignal := <-globalOSSignalCh:
			err := globalMetaDb.ActiveFile.Sync()
            globalMetaDb.Close()
			fmt.Fprintf(os.Stdout, "signal: %s", osSignal.String())
			if err == nil {
				exit(true)
			}
			fmt.Fprintf(os.Stdout, "err:%s", err.Error())
			exit(false)
		}
	}
}
