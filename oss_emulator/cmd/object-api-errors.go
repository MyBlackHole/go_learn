package emulator

import (
	"context"
	"fmt"
)

type GenericError struct {
	Bucket string
	Object string
	Err    error
}

func (e GenericError) Unwrap() error {
	return e.Err
}

type BucketNotFound GenericError

func (e BucketNotFound) Error() string {
	return "Bucket not found: " + e.Bucket
}

type BucketNameInvalid GenericError

func (e BucketNameInvalid) Error() string {
	return "Bucket name invalid: " + e.Bucket
}

type BucketNotEmpty GenericError

func (e BucketNotEmpty) Error() string {
	return "Bucket not empty: " + e.Bucket
}

type BucketExists GenericError

func (e BucketExists) Error() string {
	return "Bucket exists: " + e.Bucket
}

type StorageFull struct{}

func (e StorageFull) Error() string {
	return "Storage reached its minimum free drive threshold."
}

type BucketAlreadyOwnedByYou GenericError

func (e BucketAlreadyOwnedByYou) Error() string {
	return "Bucket already owned by you: " + e.Bucket
}

type InvalidObjectState GenericError

func (e InvalidObjectState) Error() string {
	return "The operation is not valid for the current state of the object " + e.Bucket + "/" + e.Object
}

type PreConditionFailed struct{}

func (e PreConditionFailed) Error() string {
	return "At least one of the pre-conditions you specified did not hold"
}

type BucketQuotaExceeded GenericError

func (e BucketQuotaExceeded) Error() string {
	return "Bucket quota exceeded for bucket: " + e.Bucket
}

type OperationTimedOut struct{}

func (e OperationTimedOut) Error() string {
	return "Operation timed out"
}

type ObjectNameTooLong GenericError

func (e ObjectNameTooLong) Error() string {
	return "Object name too long: " + e.Bucket + "/" + e.Object
}

type BucketAlreadyExists GenericError

func (e BucketAlreadyExists) Error() string {
	return "The requested bucket name is not available. The bucket namespace is shared by all users of the system. Please select a different name and try again."
}

type SlowDown struct{}

func (e SlowDown) Error() string {
	return "Please reduce your request rate"
}

type PrefixAccessDenied GenericError

func (e PrefixAccessDenied) Error() string {
	return "Prefix access is denied: " + e.Bucket + SlashSeparator + e.Object
}

type ObjectExistsAsDirectory GenericError

func (e ObjectExistsAsDirectory) Error() string {
	return "Object exists on : " + e.Bucket + " as directory " + e.Object
}

type VersionNotFound GenericError

func (e VersionNotFound) Error() string {
	return "Version not found: " + e.Bucket + "/" + e.Object
}

type BackendDown struct {
	Err string
}

func (e BackendDown) Error() string {
	return e.Err
}

type ObjectAlreadyExists GenericError

func (e ObjectAlreadyExists) Error() string {
	return "Object: " + e.Bucket + "/" + e.Object + " already exists"
}

type MethodNotAllowed GenericError

func (e MethodNotAllowed) Error() string {
	return "Method not allowed: " + e.Bucket + "/" + e.Object
}


type ObjectNotFound GenericError

func (e ObjectNotFound) Error() string {
	return "Object not found: " + e.Bucket + "/" + e.Object
}

type InsufficientWriteQuorum GenericError

func (e InsufficientWriteQuorum) Error() string {
	return "Storage resources are insufficient for the write operation " + e.Bucket + "/" + e.Object
}

type InvalidPart struct {
	PartNumber int
	ExpETag    string
	GotETag    string
}

func (e InvalidPart) Error() string {
	return fmt.Sprintf("Specified part could not be found. PartNumber %d, Expected %s, got %s",
		e.PartNumber, e.ExpETag, e.GotETag)
}

type InvalidUploadID struct {
	Bucket   string
	Object   string
	UploadID string
}

func (e InvalidUploadID) Error() string {
	return "Invalid upload id " + e.UploadID
}

type ObjectNamePrefixAsSlash GenericError

func (e ObjectNamePrefixAsSlash) Error() string {
	return "Object name contains forward slash as prefix: " + e.Bucket + "/" + e.Object
}

type ObjectNameInvalid GenericError

func (e ObjectNameInvalid) Error() string {
	return "Object name invalid: " + e.Bucket + "/" + e.Object
}

func toObjectErr(err error, params ...string) error {
	if err == nil {
		return nil
	}

	err = unwrapAll(err)

	if err == context.Canceled {
		return context.Canceled
	}

	switch err.Error() {
	case errVolumeNotFound.Error():
		apiErr := BucketNotFound{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		return apiErr
	case errVolumeNotEmpty.Error():
		apiErr := BucketNotEmpty{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		return apiErr
	case errVolumeExists.Error():
		apiErr := BucketExists{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		return apiErr
	case errDiskFull.Error():
		return StorageFull{}
	case errTooManyOpenFiles.Error():
		return SlowDown{}
	case errFileAccessDenied.Error():
		apiErr := PrefixAccessDenied{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		return apiErr
	case errIsNotRegular.Error():
		apiErr := ObjectExistsAsDirectory{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		return apiErr
	case errFileVersionNotFound.Error():
		apiErr := VersionNotFound{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		return apiErr
	case errMethodNotAllowed.Error():
		apiErr := MethodNotAllowed{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		return apiErr
	case errFileNotFound.Error():
		apiErr := ObjectNotFound{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		return apiErr
	case errUploadIDNotFound.Error():
		apiErr := InvalidUploadID{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		if len(params) >= 3 {
			apiErr.UploadID = params[2]
		}
		return apiErr
	case errFileNameTooLong.Error():
		apiErr := ObjectNameInvalid{}
		if len(params) >= 1 {
			apiErr.Bucket = params[0]
		}
		if len(params) >= 2 {
			apiErr.Object = decodeDirObject(params[1])
		}
		return apiErr
	}
	return err
}
