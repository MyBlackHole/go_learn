package emulator

import (
    "context"
    "io"
)

type StorageAPI interface {
    MakeVol(ctx context.Context, volume string) error
    StatVol(ctx context.Context, volume string) (vol VolInfo, err error)
	CreateFile(ctx context.Context, volume, path string, size int64, reader io.Reader) error
	AppendFile(ctx context.Context, volume string, path string, buf []byte) (err error)
    WriteMetadata(ctx context.Context, volume, path string, fi FileInfo) error
    ReadMetadata(ctx context.Context, volume, path string) (FileInfo, error)
    ListVols(ctx context.Context) (volsInfo []VolInfo, err error)
    WalkDir(ctx context.Context, opts WalkDirOptions) (fis []FileInfo, err error)
}
