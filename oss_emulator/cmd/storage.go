package emulator

import (
	"context"
	"errors"
	"io"
	"os"
	pathutil "path"
	"path/filepath"
	"sync"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/ncw/directio"
	"github.com/nutsdb/nutsdb"
)

const (
	BlockSizeSmall       = 32 * humanize.KiByte
	BlockSizeLarge       = 2 * humanize.MiByte
	BlockSizeReallyLarge = 4 * humanize.MiByte

	smallFileThreshold   = 128 * humanize.KiByte // Optimized for NVMe/SSDs
	largestFileThreshold = 64 * humanize.MiByte  // Optimized for HDDs

	writeMode = 0x1000 // O_DSYNC

	readMode = os.O_RDONLY | 0x40000 // O_NOATIME
)

const DirectioAlignSize = 4096

const metaDataReadDefault = 4 << 10

func AlignedBlock(blockSize int) []byte {
	return directio.AlignedBlock(blockSize)
}

var (
	ODirectPoolXLarge = sync.Pool{
		New: func() interface{} {
			b := AlignedBlock(BlockSizeReallyLarge)
			return &b
		},
	}
	ODirectPoolLarge = sync.Pool{
		New: func() interface{} {
			b := AlignedBlock(BlockSizeLarge)
			return &b
		},
	}
	ODirectPoolSmall = sync.Pool{
		New: func() interface{} {
			b := AlignedBlock(BlockSizeSmall)
			return &b
		},
	}
)

var metaDataPool = sync.Pool{New: func() interface{} { return make([]byte, 0, metaDataReadDefault) }}

func metaDataPoolGet() []byte {
	return metaDataPool.Get().([]byte)[:0]
}

func metaDataPoolPut(buf []byte) {
	if cap(buf) >= metaDataReadDefault && cap(buf) < metaDataReadDefault*4 {
		metaDataPool.Put(buf)
	}
}

type Storage struct {
	Disk string
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
	if !isValidVolname(volume) {
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

func (s *Storage) ListVols(ctx context.Context) (volsInfo []VolInfo, err error) {
	return listVols(ctx, s.Disk)
}

func listVols(ctx context.Context, dirPath string) ([]VolInfo, error) {
	if err := checkPathLength(dirPath); err != nil {
		return nil, err
	}
	entries, err := readDir(dirPath)
	if err != nil {
		if errors.Is(err, errFileAccessDenied) {
			return nil, errDiskAccessDenied
		} else if errors.Is(err, errFileNotFound) {
			return nil, errDiskNotFound
		}
		return nil, err
	}
	volsInfo := make([]VolInfo, 0, len(entries))
	for _, entry := range entries {
		if !HasSuffix(entry, SlashSeparator) || !isValidVolname(pathutil.Clean(entry)) {
			continue
		}
		volsInfo = append(volsInfo, VolInfo{
			Name: pathutil.Clean(entry),
		})
	}
	return volsInfo, nil
}

func (s *Storage) CreateFile(ctx context.Context, volume, path string, fileSize int64, r io.Reader) (written int64, err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return 0, err
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return 0, err
	}

	return s.writeAllDirect(ctx, filePath, fileSize, r, os.O_CREATE|os.O_WRONLY|os.O_EXCL)
}

func (s *Storage) writeAllDirect(ctx context.Context, filePath string, fileSize int64, r io.Reader, flags int) (written int64, err error) {
	// if contextCanceled(ctx) {
	// 	return 0, ctx.Err()
	// }

	parentFilePath := pathutil.Dir(filePath)
	if err = mkdirAll(parentFilePath, 0o777, s.Disk); err != nil {
		return 0, osErrToFileErr(err)
	}

	var w *os.File
	w, err = OpenFile(filePath, flags, 0o666)
	if err != nil {
		return 0, osErrToFileErr(err)
	}

    defer w.Close()

	var bufp *[]byte
	switch {
	case fileSize > 0 && fileSize >= largestFileThreshold:
		bufp = ODirectPoolXLarge.Get().(*[]byte)
		defer ODirectPoolXLarge.Put(bufp)
	case fileSize <= smallFileThreshold:
		bufp = ODirectPoolSmall.Get().(*[]byte)
		defer ODirectPoolSmall.Put(bufp)
	default:
		bufp = ODirectPoolLarge.Get().(*[]byte)
		defer ODirectPoolLarge.Put(bufp)
	}

	written, err = io.CopyBuffer(w, r, *bufp)
	if err != nil {
		w.Close()
		return 0, err
	}

	// if written < fileSize && fileSize >= 0 {
	// 	w.Close()
	// 	return errLessData
	// } else if written > fileSize && fileSize >= 0 {
	// 	w.Close()
	// 	return errMoreData
	// }

	if err = Fdatasync(w); err != nil {
		return 0, err
	}
	return
}

// func (s *Storage) writeAll(ctx context.Context, volume string, path string, b []byte, sync bool) (err error) {
// 	if contextCanceled(ctx) {
// 		return ctx.Err()
// 	}

// 	volumeDir, err := s.getVolDir(volume)
// 	if err != nil {
// 		return err
// 	}

// 	filePath := pathJoin(volumeDir, path)
// 	if err = checkPathLength(filePath); err != nil {
// 		return err
// 	}

// 	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC

// 	var w *os.File
// 	if sync {
// 		if len(b) > DirectioAlignSize {
// 			r := bytes.NewReader(b)
// 			return s.writeAllDirect(ctx, filePath, r.Size(), r, flags)
// 		}
// 		w, err = s.openFileSync(filePath, flags)
// 	} else {
// 		w, err = s.openFile(filePath, flags)
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	n, err := w.Write(b)
// 	if err != nil {
// 		w.Close()
// 		return err
// 	}

// 	if n != len(b) {
// 		w.Close()
// 		return io.ErrShortWrite
// 	}

// 	return w.Close()
// }

func (s *Storage) openFileSync(filePath string, mode int) (f *os.File, err error) {
	return s.openFile(filePath, mode|writeMode)
}

func (s *Storage) openFile(filePath string, mode int) (f *os.File, err error) {
	if err = mkdirAll(pathutil.Dir(filePath), 0o777, s.Disk); err != nil {
		return nil, osErrToFileErr(err)
	}

	w, err := OpenFile(filePath, mode, 0o666)
	if err != nil {
		switch {
		case isSysErrIsDir(err):
			return nil, errIsNotRegular
		case osIsPermission(err):
			return nil, errFileAccessDenied
		case isSysErrNotDir(err):
			return nil, errFileAccessDenied
		case isSysErrIO(err):
			return nil, errFaultyDisk
		case isSysErrTooManyFiles(err):
			return nil, errTooManyOpenFiles
		default:
			return nil, err
		}
	}

	return w, nil
}

func (s *Storage) AppendFile(ctx context.Context, volume string, path string, appendFileSize int64, r io.Reader) (written int64,  err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return
	}

	var w *os.File
	w, err = s.openFileSync(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY)
	if err != nil {
		return
	}
	defer w.Close()

	var bufp *[]byte
	bufp = ODirectPoolLarge.Get().(*[]byte)
	defer ODirectPoolLarge.Put(bufp)

	written, err = io.CopyBuffer(w, r, *bufp)
	if err != nil {
		return
	}

	// if written < appendFileSize && appendFileSize >= 0 {
	// 	return errLessData
	// } else if written > appendFileSize && appendFileSize >= 0 {
	// 	return errMoreData
	// }

	if err = Fdatasync(w); err != nil {
		return
	}

	return
}

func (s *Storage) WriteMetadata(ctx context.Context, volume, path string, fi FileInfo) (err error) {
	buf, err := fi.MarshalMsg(metaDataPoolGet())
	defer metaDataPoolPut(buf)
	if err != nil {
		return err
	}

	err = globalMetaDb.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(path)
			val := buf
			return tx.Put(volume, key, val, 0)
		})
	return
}

func (s *Storage) ReadAll(ctx context.Context, volume string, path string) (buf []byte, err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return nil, err
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return nil, err
	}

	buf, _, err = s.readAllData(ctx, volume, volumeDir, filePath, false)
	return buf, err
}

func (s *Storage) readAllData(ctx context.Context, volume, volumeDir string, filePath string, discard bool) (buf []byte, dmTime time.Time, err error) {
	if filePath == "" {
		return nil, dmTime, errFileNotFound
	}

	if contextCanceled(ctx) {
		return nil, time.Time{}, ctx.Err()
	}

	f, err := OpenFile(filePath, readMode, 0o666)
	if err != nil {
		switch {
		case osIsNotExist(err):
			if err = Access(volumeDir); err != nil && osIsNotExist(err) {
				return nil, dmTime, errVolumeNotFound
			}
			return nil, dmTime, errFileNotFound
		case osIsPermission(err):
			return nil, dmTime, errFileAccessDenied
		case isSysErrNotDir(err) || isSysErrIsDir(err):
			return nil, dmTime, errFileNotFound
		case isSysErrHandleInvalid(err):
			return nil, dmTime, errFileNotFound
		case isSysErrIO(err):
			return nil, dmTime, errFaultyDisk
		case isSysErrTooManyFiles(err):
			return nil, dmTime, errTooManyOpenFiles
		case isSysErrInvalidArg(err):
			st, _ := Lstat(filePath)
			if st != nil && st.IsDir() {
				return nil, dmTime, errFileNotFound
			}
			return nil, dmTime, errUnsupportedDisk
		}
		return nil, dmTime, err
	}

	if discard {
		defer Fdatasync(f)
	}

	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		buf, err = io.ReadAll(f)
		return buf, dmTime, osErrToFileErr(err)
	}
	if stat.IsDir() {
		return nil, dmTime, errFileNotFound
	}

	sz := stat.Size()
	if sz <= metaDataReadDefault {
		buf = metaDataPoolGet()
		buf = buf[:sz]
	} else {
		buf = make([]byte, sz)
	}

	_, err = io.ReadFull(f, buf)
	if err != nil {
		return
	}

	return buf, stat.ModTime().UTC(), osErrToFileErr(err)
}

func (s *Storage) ReadMetadata(ctx context.Context, volume, path string) (fi FileInfo, err error) {
	var buf []byte
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return
	}

	err = globalMetaDb.View(
		func(tx *nutsdb.Tx) error {
			key := []byte(path)
			buf, err = tx.Get(volume, key)
			return nil
		})

	if err != nil {
		return
	}

	if len(buf) <= 0 {
        err = errFileNotFound
		return
	}

	_, err = fi.UnmarshalMsg(buf)
	if err != nil {
		return
	}
	return
}

func isValidVolname(volname string) bool {
	return len(volname) >= 3
}

func (s *Storage) WalkDir(ctx context.Context, opts WalkDirOptions) (fis []FileInfo, err error) {
	volumeDir, err := s.getVolDir(opts.Bucket)
	if err != nil {
		return
	}

	if err = Access(volumeDir); err != nil {
		if osIsNotExist(err) {
			err = errVolumeNotFound
			return
		} else if isSysErrIO(err) {
			err = errFaultyDisk
			return
		}
		return
	}
	_, err = s.ListDir(ctx, opts.Bucket, opts.BaseDir, -1)
	if err != nil {
		if opts.ReportNotFound && err == errFileNotFound {
			err = errFileNotFound
		} else {
			err = nil
		}
		return
	}
	return
}

func (s *Storage) ListDir(ctx context.Context, volume, dirPath string, count int) (entries []string, err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return nil, err
	}

	dirPathAbs := pathJoin(volumeDir, dirPath)
	if count > 0 {
		entries, err = readDirN(dirPathAbs, count)
	} else {
		entries, err = readDir(dirPathAbs)
	}
	if err != nil {
		if err == errFileNotFound {
			if ierr := Access(volumeDir); ierr != nil {
				if osIsNotExist(ierr) {
					return nil, errVolumeNotFound
				} else if isSysErrIO(ierr) {
					return nil, errFaultyDisk
				}
			}
		}
		return nil, err
	}

	return entries, nil
}

func (s *Storage) ReadFile(ctx context.Context, volume, path string, w io.Writer) (err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return
	}

	filePath := pathJoin(volumeDir, path)

	var r *os.File
	r, err = OpenFile(filePath, readMode, 0o666)
	if err != nil {
		switch {
		case osIsNotExist(err):
			if err = Access(volumeDir); err != nil && osIsNotExist(err) {
				return errVolumeNotFound
			}
			return errFileNotFound
		case osIsPermission(err):
			return errFileAccessDenied
		case isSysErrNotDir(err) || isSysErrIsDir(err):
			return errFileNotFound
		case isSysErrHandleInvalid(err):
			return errFileNotFound
		case isSysErrIO(err):
			return errFaultyDisk
		case isSysErrTooManyFiles(err):
			return errTooManyOpenFiles
		case isSysErrInvalidArg(err):
			st, _ := Lstat(filePath)
			if st != nil && st.IsDir() {
				return errFileNotFound
			}
			return errUnsupportedDisk
		}
		return err
	}

    defer r.Close()

	var bufp *[]byte

    bufp = ODirectPoolLarge.Get().(*[]byte)
    defer ODirectPoolLarge.Put(bufp)

	_, err = io.CopyBuffer(w, r, *bufp)
	if err != nil {
		return err
	}

    return
}
