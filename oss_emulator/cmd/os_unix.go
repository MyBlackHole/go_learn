package emulator

import (
	"os"

	"golang.org/x/sys/unix"
)

func access(name string) error {
	if err := unix.Access(name, unix.F_OK); err != nil {
		return &os.PathError{Op: "lstat", Path: name, Err: err}
	}
	return nil
}
