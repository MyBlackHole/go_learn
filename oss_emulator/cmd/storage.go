package emulator

import (
	"context"
	"os"
	"path/filepath"
)

type Storage struct {
	Disk string
}

func isValidBucket(backup string) bool {
	return len(backup) >= 3
}

func getValidPath(path string) (string, error) {
	if path == "" {
		return path, errInvalidArgument
	}

	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		return path, err
	}

	fi, err := Lstat(path)
	if err != nil && !osIsNotExist(err) {
		return path, err
	}
	if osIsNotExist(err) {
		if err = mkdirAll(path, 0o777, ""); err != nil {
			return path, err
		}
	}
	if fi != nil && !fi.IsDir() {
		return path, errDiskNotDir
	}

	return path, nil
}

func newStorage(disk string) (s *Storage, err error) {
	if disk, err = getValidPath(disk); err != nil {
		return nil, err
	}

	s = &Storage{
        Disk: disk,
    }
	return s, nil
}

func checkPathLength(pathName string) error {
	if pathName == "." || pathName == ".." || pathName == slashSeparator {
		return errFileAccessDenied
	}

	var count int64
	for _, p := range pathName {
		switch p {
		case '/':
			count = 0 // Reset
		default:
			count++
			if count > 255 {
				return errFileNameTooLong
			}
		}
	} // Success.
	return nil
}

func (s *Storage) getVolDir(volume string) (string, error) {
	if volume == "" || volume == "." || volume == ".." {
		return "", errVolumeNotFound
	}
	volumeDir := pathJoin(s.Disk, volume)
	return volumeDir, nil
}

func (s *Storage) MakeVol(ctx context.Context, volume string) error {
	if !isValidBucket(volume) {
		return errInvalidArgument
	}

	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return err
	}

	if err = Access(volumeDir); err != nil {
		if osIsNotExist(err) {
			err = mkdirAll(volumeDir, 0o777, s.Disk)
		}
		if osIsPermission(err) {
			return errDiskAccessDenied
		} else if isSysErrIO(err) {
			return errFaultyDisk
		}
		return err
	}

	return nil
}

func (s *Storage) StatVol(ctx context.Context, volume string) (vol VolInfo, err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return VolInfo{}, err
	}

	var st os.FileInfo
	st, err = Lstat(volumeDir)
	if err != nil {
		switch {
		case osIsNotExist(err):
			return VolInfo{}, errVolumeNotFound
		case osIsPermission(err):
			return VolInfo{}, errDiskAccessDenied
		case isSysErrIO(err):
			return VolInfo{}, errFaultyDisk
		default:
			return VolInfo{}, err
		}
	}

	createdTime := st.ModTime()
	return VolInfo{
		Name:    volume,
		Created: createdTime,
	}, nil
}
