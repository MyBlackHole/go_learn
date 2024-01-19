package emulator

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var errInvalidArgument = errors.New("Invalid arguments specified")


type APIErrorCode int

type errorCodeMap map[APIErrorCode]APIError

func (e errorCodeMap) ToAPIErr(errCode APIErrorCode) APIError {
	return e.ToAPIErrWithErr(errCode, nil)
}

func (e errorCodeMap) ToAPIErrWithErr(errCode APIErrorCode, err error) APIError {
	apiErr, ok := e[errCode]
	if !ok {
		apiErr = e[ErrInternalError]
	}
	if err != nil {
		apiErr.Description = fmt.Sprintf("%s (%s)", apiErr.Description, err)
	}
	return apiErr
}

type APIError struct {
	Code           string
	Description    string
	HTTPStatusCode int
}

const (
	ErrNone APIErrorCode = iota
	ErrInvalidCopyDest
	ErrInternalError
	ErrServerNotInitialized
)

var noError = APIError{}

var errorCodes = errorCodeMap{
	ErrInvalidCopyDest: {
		Code:           "InvalidRequest",
		Description:    "This copy request is illegal because it is trying to copy an object to itself without changing the object's metadata, storage class, website redirect location or encryption attributes.",
		HTTPStatusCode: http.StatusBadRequest,
	},
}

func toAPIErrorCode(ctx context.Context, err error) (apiErr APIErrorCode) {
	if err == nil {
		return ErrNone
	}
	return
}

func toAPIError(ctx context.Context, err error) APIError {
	if err == nil {
		return noError
	}

	apiErr := errorCodes.ToAPIErr(toAPIErrorCode(ctx, err))
	switch apiErr.Code {
	case "NotImplemented":
		desc := fmt.Sprintf("%s (%v)", apiErr.Description, err)
		apiErr = APIError{
			Code:           apiErr.Code,
			Description:    desc,
			HTTPStatusCode: apiErr.HTTPStatusCode,
		}
	case "XMinioBackendDown":
		apiErr.Description = fmt.Sprintf("%s (%v)", apiErr.Description, err)
	case "InternalError":
		switch e := err.(type) {
		default:
			apiErr = APIError{
				Code:           apiErr.Code,
				Description:    fmt.Sprintf("%s: cause(%v), type(%s) ", apiErr.Description, err, e),
				HTTPStatusCode: apiErr.HTTPStatusCode,
			}
		}
	}

	return apiErr
}

func writeErrorResponse(ctx context.Context, w http.ResponseWriter, err APIError, reqURL *url.URL) {
	if err.HTTPStatusCode < 100 || err.HTTPStatusCode > 999 {
		fmt.Errorf("invalid WriteHeader code %v from %v", err.HTTPStatusCode, err.Code)
		err.HTTPStatusCode = http.StatusInternalServerError
	}

	errorResponse := getAPIErrorResponse(ctx, err, reqURL.Path,
		w.Header().Get(AmzRequestID), w.Header().Get(AmzRequestHostID))
	encodedErrorResponse := encodeResponse(errorResponse)
	writeResponse(w, err.HTTPStatusCode, encodedErrorResponse, mimeXML)
}

type APIErrorResponse struct {
	XMLName    xml.Name `xml:"Error" json:"-"`
	Code       string
	Message    string
	Key        string `xml:"Key,omitempty" json:"Key,omitempty"`
	BucketName string `xml:"BucketName,omitempty" json:"BucketName,omitempty"`
	Resource   string
	Region     string `xml:"Region,omitempty" json:"Region,omitempty"`
	RequestID  string `xml:"RequestId" json:"RequestId"`
	HostID     string `xml:"HostId" json:"HostId"`
}

func getAPIErrorResponse(ctx context.Context, err APIError, resource, requestID, hostID string) APIErrorResponse {
	reqInfo := GetReqInfo(ctx)
	return APIErrorResponse{
		Code:       err.Code,
		Message:    err.Description,
		BucketName: reqInfo.BucketName,
		Key:        reqInfo.ObjectName,
		Resource:   resource,
		Region:     "",
		RequestID:  requestID,
		HostID:     hostID,
	}
}
