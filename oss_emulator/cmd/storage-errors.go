package emulator

import (
	"errors"
	"os"
	"syscall"
)

type StorageErr string

var errFileNameTooLong = StorageErr("file name too long")

var errDiskNotDir = StorageErr("drive is not directory or mountpoint")

var errFileAccessDenied = StorageErr("file access denied")

var errFileNotFound = StorageErr("file not found")

var errTooManyOpenFiles = StorageErr("too many open files, please increase 'ulimit -n'")

var errFaultyDisk = StorageErr("drive is faulty")

var errDiskFull = StorageErr("drive path full")

var errVolumeNotFound = StorageErr("volume not found")

var errDiskAccessDenied = StorageErr("drive access denied")

var errLessData = StorageErr("less data available than what was requested")

var errMoreData = StorageErr("more data was sent than what was advertised")

var errIsNotRegular = StorageErr("not of regular file type")

func osIsPermission(err error) bool {
	return errors.Is(err, os.ErrPermission) || errors.Is(err, syscall.EROFS)
}

func (h StorageErr) Error() string {
	return string(h)
}

func isSysErrNotDir(err error) bool {
	return errors.Is(err, syscall.ENOTDIR)
}

func osIsNotExist(err error) bool {
	return errors.Is(err, os.ErrNotExist)
}

func osIsExist(err error) bool {
	return errors.Is(err, os.ErrExist)
}

func isSysErrPathNotFound(err error) bool {
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		var errno syscall.Errno
		if errors.As(pathErr.Err, &errno) {
			return errno == 0x03
		}
	}
	return false
}

func isSysErrIsDir(err error) bool {
	return errors.Is(err, syscall.EISDIR)
}

func isSysErrTooManyFiles(err error) bool {
	return errors.Is(err, syscall.ENFILE) || errors.Is(err, syscall.EMFILE)
}

func osErrToFileErr(err error) error {
	if err == nil {
		return nil
	}
	if osIsNotExist(err) {
		return errFileNotFound
	}
	if osIsPermission(err) {
		return errFileAccessDenied
	}
	if isSysErrNotDir(err) || isSysErrIsDir(err) {
		return errFileNotFound
	}
	if isSysErrPathNotFound(err) {
		return errFileNotFound
	}
	if isSysErrTooManyFiles(err) {
		return errTooManyOpenFiles
	}
	if isSysErrHandleInvalid(err) {
		return errFileNotFound
	}
	if isSysErrIO(err) {
		return errFaultyDisk
	}
	if isSysErrInvalidArg(err) {
		return errFileNotFound
	}
	if isSysErrNoSpace(err) {
		return errDiskFull
	}
	return err
}

func isSysErrHandleInvalid(err error) bool {
	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		var errno syscall.Errno
		if errors.As(pathErr.Err, &errno) {
			return errno == 0x6
		}
	}
	return false
}

func isSysErrIO(err error) bool {
	return errors.Is(err, syscall.EIO)
}

func isSysErrInvalidArg(err error) bool {
	return errors.Is(err, syscall.EINVAL)
}

func isSysErrNoSpace(err error) bool {
	return errors.Is(err, syscall.ENOSPC)
}
