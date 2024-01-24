package emulator

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
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

	maxObjectList = 1000

	dotdotComponent = ".."
	dotComponent    = "."
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

func getListObjectsArgs(values url.Values) (prefix, marker, delimiter string, maxkeys int, encodingType string, errCode APIErrorCode) {
	errCode = ErrNone

	if values.Get("max-keys") != "" {
		var err error
		if maxkeys, err = strconv.Atoi(values.Get("max-keys")); err != nil {
			errCode = ErrInvalidMaxKeys
			return
		}
	} else {
		maxkeys = maxObjectList
	}

	prefix = values.Get("prefix")
	marker = values.Get("marker")
	if marker == "/" {
		marker = ""
	}
	delimiter = values.Get("delimiter")
	encodingType = values.Get("encoding-type")
	return
}

func validateListObjectsArgs(prefix, marker, delimiter, encodingType string, maxKeys int) APIErrorCode {
	if maxKeys < 0 {
		return ErrInvalidMaxKeys
	}

	if encodingType != "" {
		if !strings.EqualFold(encodingType, "url") {
			return ErrInvalidEncodingMethod
		}
	}

	if !IsValidObjectPrefix(prefix) {
		return ErrInvalidObjectName
	}

	if marker != "" && !HasPrefix(marker, prefix) {
		return ErrNotImplemented
	}

	return ErrNone
}

func IsValidObjectPrefix(object string) bool {
	if hasBadPathComponent(object) {
		return false
	}
	if !utf8.ValidString(object) {
		return false
	}
	if strings.Contains(object, `//`) {
		return false
	}
	return !strings.ContainsRune(object, 0)
}

func hasBadPathComponent(path string) bool {
	path = filepath.ToSlash(strings.TrimSpace(path))
	for _, p := range strings.Split(path, SlashSeparator) {
		switch strings.TrimSpace(p) {
		case dotdotComponent:
			return true
		case dotComponent:
			return true
		}
	}
	return false
}

func HasPrefix(s string, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func generateListObjectsResponse(bucket, prefix, marker, delimiter, encodingType string, maxKeys int, resp ListObjectsInfo) ListObjectsResponse {
	contents := make([]Object, 0, len(resp.Objects))
	owner := &Owner{
		ID:          globalOssDefaultOwnerID,
		DisplayName: "oss",
	}
	data := ListObjectsResponse{}

	for _, object := range resp.Objects {
		content := Object{}
		if object.Name == "" {
			continue
		}
		content.Key = s3EncodeName(object.Name, encodingType)
		content.LastModified = ISO8601Format(object.ModTime.UTC())
		if object.ETag != "" {
			content.ETag = object.ETag
		}
		content.Size = object.Size
		content.StorageClass = object.StorageClass
		content.Type = object.Type
		content.Owner = owner
		contents = append(contents, content)
	}
	data.Name = bucket
	data.Contents = contents

	data.EncodingType = encodingType
	data.Prefix = s3EncodeName(prefix, encodingType)
	data.Marker = s3EncodeName(marker, encodingType)
	data.Delimiter = s3EncodeName(delimiter, encodingType)
	data.MaxKeys = maxKeys
	data.NextMarker = s3EncodeName(resp.NextMarker, encodingType)
	data.IsTruncated = resp.IsTruncated

	prefixes := make([]CommonPrefix, 0, len(resp.Prefixes))
	for _, prefix := range resp.Prefixes {
		prefixItem := CommonPrefix{}
		prefixItem.Prefix = s3EncodeName(prefix, encodingType)
		prefixes = append(prefixes, prefixItem)
	}
	data.CommonPrefixes = prefixes
	return data
}

type ListObjectsResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListBucketResult" json:"-"`

	Name   string
	Prefix string
	Marker string

	NextMarker string `xml:"NextMarker,omitempty"`

	MaxKeys     int
	Delimiter   string `xml:"Delimiter,omitempty"`
	IsTruncated bool

	Contents       []Object
	CommonPrefixes []CommonPrefix

	// Encoding type used to encode object keys in the response.
	EncodingType string `xml:"EncodingType,omitempty"`
}

type Object struct {
	Key          string
	LastModified string
	ETag         string
	Size         int64
	StorageClass string
	Type         string

	Owner *Owner `xml:"Owner,omitempty"`
}

type CommonPrefix struct {
	Prefix string
}
