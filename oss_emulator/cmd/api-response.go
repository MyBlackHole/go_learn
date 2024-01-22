package emulator

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
)

type mimeType string

const (
	mimeXML  mimeType = "application/xml"
	mimeNone mimeType = ""

	ContentType     = "Content-Type"
	ContentLength   = "Content-Length"
	ServerInfo      = "Server"
	AmzBucketRegion = "X-Amz-Bucket-Region"
	AcceptRanges    = "Accept-Ranges"
	Location        = "Location"

	AmzServerSideEncryption                      = "X-Amz-Server-Side-Encryption"
	AmzServerSideEncryptionKmsID                 = AmzServerSideEncryption + "-Aws-Kms-Key-Id"
	AmzServerSideEncryptionKmsContext            = AmzServerSideEncryption + "-Context"
	AmzServerSideEncryptionCustomerAlgorithm     = AmzServerSideEncryption + "-Customer-Algorithm"
	AmzServerSideEncryptionCustomerKey           = AmzServerSideEncryption + "-Customer-Key"
	AmzServerSideEncryptionCustomerKeyMD5        = AmzServerSideEncryption + "-Customer-Key-Md5"
	AmzServerSideEncryptionCopyCustomerAlgorithm = "X-Amz-Copy-Source-Server-Side-Encryption-Customer-Algorithm"
	AmzServerSideEncryptionCopyCustomerKey       = "X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key"
	AmzServerSideEncryptionCopyCustomerKeyMD5    = "X-Amz-Copy-Source-Server-Side-Encryption-Customer-Key-Md5"
	AmzMetaUnencryptedContentLength              = "X-Amz-Meta-X-Amz-Unencrypted-Content-Length"
	AmzMetaUnencryptedContentMD5                 = "X-Amz-Meta-X-Amz-Unencrypted-Content-Md5"
)

func RemoveSensitiveHeaders(h http.Header) {
	h.Del(AmzServerSideEncryptionCustomerKey)
	h.Del(AmzServerSideEncryptionCopyCustomerKey)
	h.Del(AmzMetaUnencryptedContentLength)
	h.Del(AmzMetaUnencryptedContentMD5)
}

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set(ServerInfo, "AIO OSS")

	w.Header().Set(AcceptRanges, "bytes")

	RemoveSensitiveHeaders(w.Header())
}

func writeResponse(w http.ResponseWriter, statusCode int, response []byte, mType mimeType) {
	if statusCode == 0 {
		statusCode = 200
	}
	if statusCode < 100 || statusCode > 999 {
		fmt.Errorf("invalid WriteHeader code %v", statusCode)
		statusCode = http.StatusInternalServerError
	}
	setCommonHeaders(w)
	if mType != mimeNone {
		w.Header().Set(ContentType, string(mType))
	}
	w.Header().Set(ContentLength, strconv.Itoa(len(response)))
	w.WriteHeader(statusCode)
	if response != nil {
		w.Write(response)
	}
}

func writeSuccessResponseHeadersOnly(w http.ResponseWriter) {
	writeResponse(w, http.StatusOK, nil, mimeNone)
}

func methodNotAllowedHandler(api string) func(w http.ResponseWriter, r *http.Request) {
	return errorResponseHandler
}

func errorResponseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	switch {
	default:
		writeErrorResponse(r.Context(), w, APIError{
			Code: "BadRequest",
			Description: fmt.Sprintf("An error occurred when parsing the HTTP request %s at '%s'",
				r.Method, r.URL.Path),
			HTTPStatusCode: http.StatusBadRequest,
		}, r.URL)
	}
}

func writeErrorResponseHeadersOnly(w http.ResponseWriter, err APIError) {
	w.Header().Set(xOSSIOErrCodeHeader, err.Code)
	w.Header().Set(xOSSErrDescHeader, "\""+err.Description+"\"")
	writeResponse(w, err.HTTPStatusCode, nil, mimeNone)
}

type Owner struct {
	ID          string
	DisplayName string
}

type Bucket struct {
	Name         string
	CreationDate string
}

type ListBucketsResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsResult" json:"-"`

	Owner Owner

	Buckets struct {
		Buckets []Bucket `xml:"Bucket"`
	}
}

func generateListBucketsResponse(buckets []BucketInfo) ListBucketsResponse {
	listbuckets := make([]Bucket, 0, len(buckets))
	data := ListBucketsResponse{}
	owner := Owner{
		ID:          globalOssDefaultOwnerID,
		DisplayName: "aio-oss",
	}

	for _, bucket := range buckets {
		listbuckets = append(listbuckets, Bucket{
			Name:         bucket.Name,
			CreationDate: ISO8601Format(bucket.Created.UTC()),
		})
	}

	data.Owner = owner
	data.Buckets.Buckets = listbuckets

	return data
}

func writeSuccessResponseXML(w http.ResponseWriter, response []byte) {
	writeResponse(w, http.StatusOK, response, mimeXML)
}
