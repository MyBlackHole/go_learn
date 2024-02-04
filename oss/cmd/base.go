package oss

var (
	VERSION        = "1.0.0"
	VERSION_STRING = "Emulator V1.0.0"

	STORE_ROOT_DIR        = "store"
	BUCKET_METADATA       = ".metadata_bucket_oss_aliyun_ALIBABA"
	OBJECT_METADATA       = ".metadata_object_oss_aliyun_ALIBABA"
	OBJECT_CONTENT_PREFIX = ".content_object_oss_aliyun_ALIBABA-"
	OBJECT_CONTENT        = ".content_object_oss_aliyun_ALIBABA-1"
	OBJECT_CONTENT_TWO    = ".content_object_oss_aliyun_ALIBABA-2"

	MAX_BUCKET_NUM         = 30
	BUCKET_NAME_CHAR_SET   = "1234567890abcdefghijklmnopqrstuvwxyz-"
	MAX_BUCKET_NAME_LENGTH = 63

	MAX_OBJECT_NAME_LENGTH       = 1023
	MAX_OBJECT_NAME_SLICE_LENGTH = 255
	MAX_OBJECT_FILE_SIZE         = 5.4 * 1024 * 1024 * 1024 // 5G = 5*1024*1024*1024 Bytes
	STREAM_CHUNK_SIZE            = 32 * 1024

	HOST                = "localhost"
	HOSTNAMES           = []string{"localhost", "oss.aliyun.com", "oss.localhost"}
	ALIYUN_OSS_SERVER   = "AliyunOSS"
	REQUEST_ID          = "1234567890ABCDEF12345678"
	SUBSECOND_PRECISION = 3
	XMLNS               = "http://doc.oss-cn-hangzhou.aliyuncs.com"
	OWNER_ID            = "00220120222"
	OWNER_DISPLAY_NAME  = "1390402650033798"

	PUBLIC_READ_WRITE = "public-read-write"
	PUBLIC_READ       = "public-read"
	PRIVATE           = "private"
	DEFAULT           = "default"
	OSS_ACL           = []string{PUBLIC_READ_WRITE, PUBLIC_READ, PRIVATE, DEFAULT}
	BUCKET_ACL_LIST   = []string{PUBLIC_READ_WRITE, PUBLIC_READ, PRIVATE}
	OBJECT_ACL_LIST   = []string{PUBLIC_READ_WRITE, PUBLIC_READ, PRIVATE, DEFAULT}
)

type ErrorCode struct {
	ErrorCode          string
	StatusCode         int
	Message            string

	Bucket             string
	Object             string

	// ContentType        string
	// ContentEncoding    string
	// BucketName         string
	// ContentDisposition string
	// ModifiedDate       string
	// Md5                string
	// CustomMetadata     map[string]string
	// Multipart          string
	// ContentRange       string
	// Pos                string
	// ReadLength         string
	// BytesToRead        string
	// Cmd                string
}

var (
	NO_MODIFIED = ErrorCode{
		ErrorCode:  "NoModified",
		StatusCode: 304,
		Message:    "The object has not been modified.",
	}

	BAD_REQUEST = ErrorCode{
		ErrorCode:  "BadRequest",
		StatusCode: 400,
		Message:    "The server cannot understand the request.",
	}
	TOO_MANY_BUCKETS = ErrorCode{
		ErrorCode:  "TooManyBuckets",
		StatusCode: 400,
		Message:    "Bucket number exceeds the limit.",
	}
	INVALID_BUCKET_NAME = ErrorCode{
		ErrorCode:  "InvalidBucketName",
		StatusCode: 400,
		Message:    "The bucket name is invalid.",
	}
	INVALID_OBJECT_NAME = ErrorCode{
		ErrorCode:  "InvalidObjectName",
		StatusCode: 400,
		Message:    "The object name is invalid.",
	}
	INVALID_ARGUMENT = ErrorCode{
		ErrorCode:  "InvalidArgument",
		StatusCode: 400,
		Message:    "The file size should be less than 5G.",
	}
	FILE_PART_NO_EXIST = ErrorCode{
		ErrorCode:  "FilePartNotExist",
		StatusCode: 400,
		Message:    "The file part does not exist.",
	}
	ACCESS_DENIED = ErrorCode{
		ErrorCode:  "AccessDenied",
		StatusCode: 403,
		Message:    "The access is forbidden.",
	}
	NO_SUCH_BUCKET = ErrorCode{
		ErrorCode:  "NoSuchBucket",
		StatusCode: 404,
		Message:    "The bucket does not exist.",
	}
	NO_SUCH_KEY = ErrorCode{
		ErrorCode:  "NoSuchKey",
		StatusCode: 404,
		Message:    "The specified object does not exist.",
	}
	NOT_FOUND = ErrorCode{
		ErrorCode:  "NotFound",
		StatusCode: 404,
		Message:    "The file has not been found.",
	}
	BUCKET_ALREADY_EXISTS = ErrorCode{
		ErrorCode:  "BucketAlreadyExists",
		StatusCode: 409,
		Message:    "The bucket already exists.",
	}
	BUCKET_NOT_EMPTY = ErrorCode{
		ErrorCode:  "BucketNotEmpty",
		StatusCode: 409,
		Message:    "The bucket is not empty.",
	}
	OBJECT_NOT_APPENDABLE = ErrorCode{
		ErrorCode:  "ObjectNotAppendable",
		StatusCode: 409,
		Message:    "The object is not appendable.",
	}
	MISSING_CONTENT_LENGTH = ErrorCode{
		ErrorCode:  "MissingContentLength",
		StatusCode: 411,
		Message:    "No Content-Length in request header.",
	}
	INTERNAL_ERROR = ErrorCode{
		ErrorCode:  "InternalError",
		StatusCode: 500,
		Message:    "An internal error occurs inside OSS.",
	}
	NOT_IMPLEMENTED = ErrorCode{
		ErrorCode:  "NotImplemented",
		StatusCode: 501,
		Message:    "The function is not supported yet.",
	}
)
