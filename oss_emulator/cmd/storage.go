package emulator

import (
	"bytes"
	"context"
	"github.com/dustin/go-humanize"
	"github.com/ncw/directio"
	"io"
	"os"
	pathutil "path"
	"path/filepath"
	"sync"
)

const (
	BlockSizeSmall       = 32 * humanize.KiByte
	BlockSizeLarge       = 2 * humanize.MiByte
	BlockSizeReallyLarge = 4 * humanize.MiByte

	smallFileThreshold   = 128 * humanize.KiByte // Optimized for NVMe/SSDs
	largestFileThreshold = 64 * humanize.MiByte  // Optimized for HDDs

	writeMode = 0x1000 // O_DSYNC

	oStorageFormatFile = "o.meta"
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

func (s *Storage) CreateFile(ctx context.Context, volume, path string, fileSize int64, r io.Reader) (err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return err
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return err
	}

	return s.writeAllDirect(ctx, filePath, fileSize, r, os.O_CREATE|os.O_WRONLY|os.O_EXCL)
}

func (s *Storage) writeAllDirect(ctx context.Context, filePath string, fileSize int64, r io.Reader, flags int) (err error) {
	if contextCanceled(ctx) {
		return ctx.Err()
	}

	parentFilePath := pathutil.Dir(filePath)
	if err = mkdirAll(parentFilePath, 0o777, s.Disk); err != nil {
		return osErrToFileErr(err)
	}

	var w *os.File
	w, err = OpenFile(filePath, flags, 0o666)
	if err != nil {
		return osErrToFileErr(err)
	}

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

	var written int64
	written, err = io.CopyBuffer(w, r, *bufp)
	if err != nil {
		w.Close()
		return err
	}

	if written < fileSize && fileSize >= 0 {
		w.Close()
		return errLessData
	} else if written > fileSize && fileSize >= 0 {
		w.Close()
		return errMoreData
	}

	if err = Fdatasync(w); err != nil {
		w.Close()
		return err
	}
	return w.Close()
}

func (s *Storage) WriteMetadata(ctx context.Context, volume, path string, fi FileInfo) error {
	buf, err := fi.MarshalMsg(metaDataPoolGet())
	defer metaDataPoolPut(buf)
	if err != nil {
		return err
	}

	return s.writeAll(ctx, volume, pathJoin(path, oStorageFormatFile), buf, false)
}

func (s *Storage) writeAll(ctx context.Context, volume string, path string, b []byte, sync bool) (err error) {
	if contextCanceled(ctx) {
		return ctx.Err()
	}

	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return err
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return err
	}

	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC

	var w *os.File
	if sync {
		if len(b) > DirectioAlignSize {
			r := bytes.NewReader(b)
			return s.writeAllDirect(ctx, filePath, r.Size(), r, flags)
		}
		w, err = s.openFileSync(filePath, flags)
	} else {
		w, err = s.openFile(filePath, flags)
	}
	if err != nil {
		return err
	}

	n, err := w.Write(b)
	if err != nil {
		w.Close()
		return err
	}

	if n != len(b) {
		w.Close()
		return io.ErrShortWrite
	}

	return w.Close()
}

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

func (s *Storage) AppendFile(ctx context.Context, volume string, path string, buf []byte) (err error) {
	volumeDir, err := s.getVolDir(volume)
	if err != nil {
		return err
	}

	filePath := pathJoin(volumeDir, path)
	if err = checkPathLength(filePath); err != nil {
		return err
	}

	var w *os.File
	w, err = s.openFileSync(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY)
	if err != nil {
		return err
	}
	defer w.Close()

	n, err := w.Write(buf)
	if err != nil {
		return err
	}

	if n != len(buf) {
		return io.ErrShortWrite
	}

	return nil
}
