package emulator

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

const (
	Range            = "Range"
	AmzRequestID     = "x-amz-request-id"
	AmzRequestHostID = "x-amz-id-2"
	Header           = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
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

func encodeResponseList(response interface{}) []byte {
	var buf bytes.Buffer
	buf.WriteString(Header)
	if err := xml.NewEncoder(&buf).Encode(response); err != nil {
		return nil
	}
	return buf.Bytes()
}
