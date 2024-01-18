package emulator

import (
    "context"
)

type Objects struct {
}


func (o *Objects) MakeBucket(ctx context.Context, bucket string) error {
	if !isValidVolname(bucket) {
		return errInvalidArgument
	}

    return nil
}
