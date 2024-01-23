package emulator

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// ResponseRecorder - is a wrapper to trap the http response
// status code and to record the response body
type ResponseRecorder struct {
	http.ResponseWriter
	io.ReaderFrom
	StatusCode int
	// Log body of 4xx or 5xx responses
	LogErrBody bool
	// Log body of all responses
	LogAllBody bool

	TimeToFirstByte time.Duration
	StartTime       time.Time
	// number of bytes written
	bytesWritten int
	// number of bytes of response headers written
	headerBytesWritten int
	// Internal recording buffer
	headers bytes.Buffer
	body    bytes.Buffer
	// Indicate if headers are written in the log
	headersLogged bool
}

// Hijack - hijacks the underlying connection
func (lrw *ResponseRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, ok := lrw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("response writer does not support hijacking. Type is %T", lrw.ResponseWriter)
	}
	return hj.Hijack()
}

// NewResponseRecorder - returns a wrapped response writer to trap
// http status codes for auditing purposes.
func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	rf, _ := w.(io.ReaderFrom)
	return &ResponseRecorder{
		ResponseWriter: w,
		ReaderFrom:     rf,
		StatusCode:     http.StatusOK,
		StartTime:      time.Now().UTC(),
	}
}

// ErrNotImplemented when a functionality is not implemented
var ErrNotImplementedOld = errors.New("not implemented")

// ReadFrom implements support for calling internal io.ReaderFrom implementations
// returns an error if the underlying ResponseWriter does not implement io.ReaderFrom
func (lrw *ResponseRecorder) ReadFrom(r io.Reader) (int64, error) {
	if lrw.ReaderFrom != nil {
		n, err := lrw.ReaderFrom.ReadFrom(r)
		lrw.bytesWritten += int(n)
		return n, err
	}
	return 0, ErrNotImplementedOld
}

func (lrw *ResponseRecorder) Write(p []byte) (int, error) {
	if !lrw.headersLogged {
		// We assume the response code to be '200 OK' when WriteHeader() is not called,
		// that way following Golang HTTP response behavior.
		lrw.WriteHeader(http.StatusOK)
	}
	n, err := lrw.ResponseWriter.Write(p)
	lrw.bytesWritten += n
	if lrw.TimeToFirstByte == 0 {
		lrw.TimeToFirstByte = time.Now().UTC().Sub(lrw.StartTime)
	}
	gzipped := lrw.Header().Get("Content-Encoding") == "gzip"
	if !gzipped && ((lrw.LogErrBody && lrw.StatusCode >= http.StatusBadRequest) || lrw.LogAllBody) {
		// Always logging error responses.
		lrw.body.Write(p)
	}
	if err != nil {
		return n, err
	}
	return n, err
}

// Write the headers into the given buffer
func (lrw *ResponseRecorder) writeHeaders(w io.Writer, statusCode int, headers http.Header) {
	n, _ := fmt.Fprintf(w, "%d %s\n", statusCode, http.StatusText(statusCode))
	lrw.headerBytesWritten += n
	for k, v := range headers {
		n, _ := fmt.Fprintf(w, "%s: %s\n", k, v[0])
		lrw.headerBytesWritten += n
	}
}

// blobBody returns a dummy body placeholder for blob (binary stream)
var blobBody = []byte("<BLOB>")

// gzippedBody returns a dummy body placeholder for gzipped content
var gzippedBody = []byte("<GZIP>")

// Body - Return response body.
func (lrw *ResponseRecorder) Body() []byte {
	if lrw.Header().Get("Content-Encoding") == "gzip" {
		// ... otherwise we return the <GZIP> place holder
		return gzippedBody
	}
	// If there was an error response or body logging is enabled
	// then we return the body contents
	if (lrw.LogErrBody && lrw.StatusCode >= http.StatusBadRequest) || lrw.LogAllBody {
		return lrw.body.Bytes()
	}
	// ... otherwise we return the <BLOB> place holder
	return blobBody
}

// WriteHeader - writes http status code
func (lrw *ResponseRecorder) WriteHeader(code int) {
	if !lrw.headersLogged {
		lrw.StatusCode = code
		lrw.writeHeaders(&lrw.headers, code, lrw.ResponseWriter.Header())
		lrw.headersLogged = true
		lrw.ResponseWriter.WriteHeader(code)
	}
}

// Flush - Calls the underlying Flush.
func (lrw *ResponseRecorder) Flush() {
	lrw.ResponseWriter.(http.Flusher).Flush()
}

// Size - returns  the number of bytes written
func (lrw *ResponseRecorder) Size() int {
	return lrw.bytesWritten
}

// HeaderSize - returns the number of bytes of response headers written
func (lrw *ResponseRecorder) HeaderSize() int {
	return lrw.headerBytesWritten
}
