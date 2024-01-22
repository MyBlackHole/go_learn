package emulator

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const blockSize = 8 << 10 // 8192

var (
	direntPool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte, blockSize*128)
			return &buf
		},
	}

	direntNamePool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte, blockSize)
			return &buf
		},
	}
)


type readDirOpts struct {
	count int
	followDirSymlink bool
}

func Stat(name string) (info os.FileInfo, err error) {
	return os.Stat(name)
}

func access(name string) error {
	if err := unix.Access(name, unix.F_OK); err != nil {
		return &os.PathError{Op: "lstat", Path: name, Err: err}
	}
	return nil
}

func readDir(dirPath string) (entries []string, err error) {
	return readDirWithOpts(dirPath, readDirOpts{count: -1})
}

func readDirWithOpts(dirPath string, opts readDirOpts) (entries []string, err error) {
	fd, err := openFileWithFD(dirPath, readMode, 0o666)
	if err != nil {
		if !osIsPermission(err) {
			return nil, osErrToFileErr(err)
		}
		fd, err = openFileWithFD(dirPath, os.O_RDONLY, 0o666)
		if err != nil {
			return nil, osErrToFileErr(err)
		}
	}
	defer syscall.Close(fd)

	bufp := direntPool.Get().(*[]byte)
	defer direntPool.Put(bufp)
	buf := *bufp

	nameTmp := direntNamePool.Get().(*[]byte)
	defer direntNamePool.Put(nameTmp)
	tmp := *nameTmp

	boff := 0
	nbuf := 0

	count := opts.count

	for count != 0 {
		if boff >= nbuf {
			boff = 0
			nbuf, err = syscall.ReadDirent(fd, buf)
			if err != nil {
				if isSysErrNotDir(err) {
					return nil, errFileNotFound
				}
				return nil, osErrToFileErr(err)
			}
			if nbuf <= 0 {
				break
			}
		}
		consumed, name, typ, err := parseDirEnt(buf[boff:nbuf])
		if err != nil {
			return nil, err
		}
		boff += consumed
		if len(name) == 0 || bytes.Equal(name, []byte{'.'}) || bytes.Equal(name, []byte{'.', '.'}) {
			continue
		}

		if typ == unexpectedFileMode || typ&os.ModeSymlink == os.ModeSymlink {
			fi, err := Stat(pathJoin(dirPath, string(name)))
			if err != nil {
				if osIsNotExist(err) || isSysErrPathNotFound(err) ||
					isSysErrTooManySymlinks(err) {
					continue
				}
				return nil, err
			}

			if !opts.followDirSymlink && typ&os.ModeSymlink == os.ModeSymlink && fi.IsDir() {
				continue
			}

			typ = fi.Mode() & os.ModeType
		}

		var nameStr string
		if typ.IsRegular() {
			nameStr = string(name)
		} else if typ.IsDir() {
			tmp = tmp[:len(name)+1]
			copy(tmp, name)
			tmp[len(tmp)-1] = '/'
			nameStr = string(tmp)
		}

		count--
		entries = append(entries, nameStr)
	}

	return
}

func openFileWithFD(name string, flag int, perm os.FileMode) (fd int, err error) {
	var e error
	fd, e = syscall.Open(name, flag|syscall.O_CLOEXEC, uint32(perm))
	if e != nil {
		return -1, &os.PathError{Op: "open", Path: name, Err: e}
	}
	return fd, nil
}

const unexpectedFileMode os.FileMode = os.ModeNamedPipe | os.ModeSocket | os.ModeDevice

func parseDirEnt(buf []byte) (consumed int, name []byte, typ os.FileMode, err error) {
	dirent := (*syscall.Dirent)(unsafe.Pointer(&buf[0]))
	if v := unsafe.Offsetof(dirent.Reclen) + unsafe.Sizeof(dirent.Reclen); uintptr(len(buf)) < v {
		return consumed, nil, typ, fmt.Errorf("buf size of %d smaller than dirent header size %d", len(buf), v)
	}
	if len(buf) < int(dirent.Reclen) {
		return consumed, nil, typ, fmt.Errorf("buf size %d < record length %d", len(buf), dirent.Reclen)
	}
	consumed = int(dirent.Reclen)
	if direntInode(dirent) == 0 {
		return
	}
	switch dirent.Type {
	case syscall.DT_REG:
		typ = 0
	case syscall.DT_DIR:
		typ = os.ModeDir
	case syscall.DT_LNK:
		typ = os.ModeSymlink
	default:
		typ = unexpectedFileMode
	}

	nameBuf := (*[unsafe.Sizeof(dirent.Name)]byte)(unsafe.Pointer(&dirent.Name[0]))
	nameLen, err := direntNamlen(dirent)
	if err != nil {
		return consumed, nil, typ, err
	}

	return consumed, nameBuf[:nameLen], typ, nil
}

func direntInode(dirent *syscall.Dirent) uint64 {
	return dirent.Ino
}

func direntNamlen(dirent *syscall.Dirent) (uint64, error) {
	const fixedHdr = uint16(unsafe.Offsetof(syscall.Dirent{}.Name))
	nameBuf := (*[unsafe.Sizeof(dirent.Name)]byte)(unsafe.Pointer(&dirent.Name[0]))
	const nameBufLen = uint16(len(nameBuf))
	limit := dirent.Reclen - fixedHdr
	if limit > nameBufLen {
		limit = nameBufLen
	}
	nameLen := bytes.IndexByte(nameBuf[:limit], 0)
	if nameLen < 0 {
		return 0, fmt.Errorf("failed to find terminating 0 byte in dirent")
	}
	return uint64(nameLen), nil
}
