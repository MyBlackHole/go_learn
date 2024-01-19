package emulator

import (
	"os"
	"strings"
    "syscall"
)


func Lstat(name string) (info os.FileInfo, err error) {
	return os.Lstat(name)
}

func mkdirAll(dirPath string, mode os.FileMode, baseDir string) (err error) {
	if dirPath == "" {
		return errInvalidArgument
	}

	if err = checkPathLength(dirPath); err != nil {
		return err
	}

	if err = reliableMkdirAll(dirPath, mode, baseDir); err != nil {
		if isSysErrNotDir(err) {
			return errFileAccessDenied
		} else if isSysErrPathNotFound(err) {
			return errFileAccessDenied
		}
		return osErrToFileErr(err)
	}

	return nil
}

func reliableMkdirAll(dirPath string, mode os.FileMode, baseDir string) (err error) {
	i := 0
	for {
		if err = osMkdirAll(dirPath, mode, baseDir); err != nil {
			if osIsNotExist(err) && i == 0 {
				i++
				continue
			}
		}
		break
	}
	return err
}

func osMkdirAll(dirPath string, perm os.FileMode, baseDir string) error {
	if baseDir != "" {
		if strings.HasPrefix(baseDir, dirPath) {
			return nil
		}
	}

	// Slow path: make sure parent exists and then call Mkdir for path.
	i := len(dirPath)
	for i > 0 && os.IsPathSeparator(dirPath[i-1]) { // Skip trailing path separator.
		i--
	}

	j := i
	for j > 0 && !os.IsPathSeparator(dirPath[j-1]) { // Scan backward over element.
		j--
	}

	if j > 1 {
		// Create parent.
		if err := osMkdirAll(dirPath[:j-1], perm, baseDir); err != nil {
			return err
		}
	}

	// Parent now exists; invoke Mkdir and use its result.
	if err := Mkdir(dirPath, perm); err != nil {
		if osIsExist(err) {
			return nil
		}
		return err
	}

	return nil
}

func Mkdir(dirPath string, mode os.FileMode) (err error) {
	return os.Mkdir(dirPath, mode)
}

func Access(name string) (err error) {
	return access(name)
}

func OpenFile(name string, flag int, perm os.FileMode) (f *os.File, err error) {
	return os.OpenFile(name, flag, perm)
}

func Fdatasync(f *os.File) error {
	return syscall.Fdatasync(int(f.Fd()))
}
