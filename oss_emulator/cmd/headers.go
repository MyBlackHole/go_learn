package emulator

import (
	"net/http"
	"bytes"
	"encoding/xml"
)

const (
	AmzRequestID     = "x-amz-request-id"
	AmzRequestHostID = "x-amz-id-2"
)

func getHostName(r *http.Request) (hostName string) {
	hostName = r.Host
	return
}

func encodeResponse(response interface{}) []byte {
	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	if err := xml.NewEncoder(&buf).Encode(response); err != nil {
		return nil
	}
	return buf.Bytes()
}

