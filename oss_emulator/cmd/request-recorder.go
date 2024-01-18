package emulator

import (
	"bytes"
	"io"
)

// RequestRecorder - records the
// of a given io.Reader
type RequestRecorder struct {
	// Data source to record
	io.Reader
	// Response body should be logged
	LogBody bool

	// internal recording buffer
	buf bytes.Buffer
	// total bytes read including header size
	bytesRead int
}

// Close is a no operation closer
func (r *RequestRecorder) Close() error {
	// no-op
	return nil
}

// Read reads from the internal reader and counts/save the body in the memory
func (r *RequestRecorder) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.bytesRead += n

	if r.LogBody {
		r.buf.Write(p[:n])
	}
	if err != nil {
		return n, err
	}
	return n, err
}

// Size returns the body size if the currently read object
func (r *RequestRecorder) Size() int {
	return r.bytesRead
}

// Data returns the bytes that were recorded.
func (r *RequestRecorder) Data() []byte {
	// If body logging is enabled then we return the actual body
	if r.LogBody {
		return r.buf.Bytes()
	}
	// ... otherwise we return <BLOB> placeholder
	return blobBody
}
