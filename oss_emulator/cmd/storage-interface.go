package emulator

import (
    "context"
)

type StorageAPI interface {
    MakeVol(ctx context.Context, volume string) error
    StatVol(ctx context.Context, volume string) (vol VolInfo, err error)
}
