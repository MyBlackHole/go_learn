package oss

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var (
	PUT_BUCKET           = "PUT_BUCKET"
	PUT_BUCKET_ACL       = "PUT_BUCKET_ACL"
	PUT_BUCKET_LOGGING   = "PUT_BUCKET_LOGGING"
	PUT_BUCKET_REFERER   = "PUT_BUCKET_REFERER"
	PUT_BUCKET_WEBSITE   = "PUT_BUCKET_WEBSITE"
	PUT_BUCKET_LIFECYCLE = "PUT_BUCKET_LIFECYCLE"

	PUT_OBJECT           = "PUT_OBJECT"
	PUT_OBJECT_ACL       = "PUT_OBJECT_ACL"
	PUT_SYMLINK          = "PUT_SYMLINK"
	PUT_UPLOAD_PART      = "PUT_UPLOAD_PART"
	PUT_UPLOAD_PART_COPY = "PUT_UPLOAD_PART_COPY"
	PUT_COPY_OBJECT      = "PUT_COPY_OBJECT"

	LIST_BUCKETS         = "LIST_BUCKETS"
	GET_BUCKET           = "GET_BUCKET"
	GET_BUCKET_ACL       = "GET_BUCKET_ACL"
	GET_BUCKET_INFO      = "GET_BUCKET_INFO"
	GET_BUCKET_LOCATION  = "GET_BUCKET_LOCATION"
	GET_BUCKET_LOGGING   = "GET_BUCKET_LOGGING"
	GET_BUCKET_REFERER   = "GET_BUCKET_REFERER"
	GET_BUCKET_WEBSITE   = "GET_BUCKET_WEBSITE"
	GET_BUCKET_LIFECYCLE = "GET_BUCKET_LIFECYCLE"

	GET_OBJECT      = "GET_OBJECT"
	GET_OBJECT_ACL  = "GET_OBJECT_ACL"
	GET_OBJECT_META = "GET_OBJECT_META"
	GET_SYMLINK     = "GET_SYMLINK"

	GET_LIST_MULTIPART_UPLOADS = "GET_LIST_MULTIPART_UPLOADS"
	GET_LIST_PARTS             = "GET_LIST_PARTS"

	HEAD_OBJECT = "HEAD_OBJECT"

	DELETE_BUCKET                 = "DELETE_BUCKET"
	DELETE_BUCKET_LOGGING         = "DELETE_BUCKET_LOGGING"
	DELETE_BUCKET_WEBSITE         = "DELETE_BUCKET_WEBSITE"
	DELETE_BUCKET_LIFECYCLE       = "DELETE_BUCKET_LIFECYCLE"
	DELETE_ABORT_MULTIPART_UPLOAD = "DELETE_ABORT_MULTIPART_UPLOAD"

	DELETE_OBJECT           = "DELETE_OBJECT"
	DELETE_MULTIPLE_OBJECTS = "DELETE_MULTIPLE_OBJECTS"

	POST_RESTORE_OBJECT = "POST_RESTORE_OBJECT"

	POST_APPEND_OBJECT             = "POST_APPEND_OBJECT"
	POST_INIT_MULTIPART_UPLOAD     = "POST_INIT_MULTIPART_UPLOAD"
	POST_COMPLETE_MULTIPART_UPLOAD = "POST_COMPLETE_MULTIPART_UPLOAD"
	POST_OBJECT                    = "POST_OBJECT"
	POST_ELSE                      = "POST_ELSE"

	REQUEST_ERROR = "REQUEST_ERROR"
)

type OssRequest struct {
	*http.Request
	Query        url.Values
	Header       http.Header
	Method       string
	Cmd          string
	Bucket       string
	Object       string
	SrcBucket    string
	SrcObject    string
	Path         string
	CreationDate string
	// @path = @request.path
	// @path_length = @request.path.size
	// @cmd ||= ''
	// @bucket ||= ''
	// @object ||= ''
	// @src_bucket ||= ''
	// @src_object ||= ''
}

func NewOssRequest(request *http.Request) *OssRequest {
	return &OssRequest{
		Request: request,
		Query:   request.URL.Query(),
		Header:  request.Header,
		Method:  request.Method,
		Path:    request.URL.Path,
	}
}

func (ossRequest *OssRequest) ValidateRequest() {
	if ossRequest.Header == nil {
		return
	}

	if ossRequest.Header.Get("expect") == "" {
		return
	}
	// TODO 啥
	// @request.continue if @request.header['expect'].first=='100-continue'
}

// func (ossRequest *OssRequest) Parse() {
// 	// host := ossRequest.Host
// 	method := ossRequest.Method

// 	switch method {
// 	case "GET", "HEAD":
// 		ossRequest.ParseGet()
// 	case "PUT":
// 		ossRequest.ParsePut()
// 	case "DELETE":
// 		ossRequest.ParseDelete()
// 	case "POST":
// 		ossRequest.ParsePost()
// 	default:
// 		panic(method)
// 	}

// 	// validate_request()
// 	// inspect_request()
// }

func (ossRequest *OssRequest) ParseGet() {
	// TODO is_path_style 没有处理
	if ossRequest.Path == "/" {
		switch {
		case ossRequest.Query.Has("uploads"):
			ossRequest.Cmd = GET_LIST_MULTIPART_UPLOADS
		case ossRequest.Query.Has("logging"):
			ossRequest.Cmd = GET_BUCKET_LOGGING
		case ossRequest.Query.Has("website"):
			ossRequest.Cmd = GET_BUCKET_WEBSITE
		case ossRequest.Query.Has("referer"):
			ossRequest.Cmd = GET_BUCKET_REFERER
		case ossRequest.Query.Has("lifecycle"):
			ossRequest.Cmd = GET_BUCKET_LIFECYCLE
		default:
			ossRequest.Cmd = LIST_BUCKETS
		}
	} else {
		// TODO is_path_style
		elems := strings.Split(ossRequest.Path[1:], "/")
		ossRequest.Bucket = elems[0]

		if len(elems) < 2 || (len(elems) == 2 && elems[1] == "") {
			switch {
			case ossRequest.Query.Has("acl"):
				ossRequest.Cmd = GET_BUCKET_ACL
			case ossRequest.Query.Has("location"):
				ossRequest.Cmd = GET_BUCKET_LOCATION
			case ossRequest.Query.Has("bucketInfo"):
				ossRequest.Cmd = GET_BUCKET_INFO
			default:
				ossRequest.Cmd = GET_BUCKET
			}
		} else {
			switch {
			case ossRequest.Query.Has("acl"):
				ossRequest.Cmd = GET_BUCKET_ACL
			case ossRequest.Query.Has("objectMeta"):
				ossRequest.Cmd = GET_OBJECT_META
			case ossRequest.Query.Has("symlink"):
				ossRequest.Cmd = GET_SYMLINK
			case ossRequest.Query.Has("uploadId"):
				ossRequest.Cmd = GET_LIST_PARTS
			default:
                if ossRequest.Method == "HEAD" {
                    ossRequest.Cmd = HEAD_OBJECT
                } else {
                    ossRequest.Cmd = GET_OBJECT
                }
			}
		}
		ossRequest.Object = strings.Join(elems[1:], "/")
	}
}

func (ossRequest *OssRequest) ParseDelete() {
	// TODO is_path_style
	if ossRequest.Path == "/" {
		ossRequest.Cmd = REQUEST_ERROR
	} else {
		// TODO is_path_style
		elems := strings.Split(ossRequest.Path[1:], "/")
		ossRequest.Bucket = elems[0]
		switch len(elems) {
		case 0:
			panic("elems.size = 0")
		case 1:
			switch {
			case ossRequest.Query.Has("logging"):
				ossRequest.Cmd = DELETE_BUCKET_LOGGING
			case ossRequest.Query.Has("website"):
				ossRequest.Cmd = DELETE_BUCKET_WEBSITE
			case ossRequest.Query.Has("lifecycle"):
				ossRequest.Cmd = DELETE_BUCKET_LIFECYCLE
			default:
				ossRequest.Cmd = DELETE_BUCKET
			}
		default:
			if ossRequest.Query.Has("logging") {
				ossRequest.Cmd = DELETE_ABORT_MULTIPART_UPLOAD
			} else {
				ossRequest.Cmd = DELETE_OBJECT
			}
			ossRequest.Object = strings.Join(elems[1:], "/")
		}
	}
}

func (ossRequest *OssRequest) ParsePost() {
	elems := strings.Split(ossRequest.Path[1:], "/")
	ossRequest.Bucket = elems[0]
	// TODO elems.size > 2 ?
	ossRequest.Object = strings.Join(elems[1:], "/")
	switch {
	case ossRequest.Query.Has("uploads"):
		ossRequest.Cmd = POST_INIT_MULTIPART_UPLOAD
	case ossRequest.Query.Has("append"):
		ossRequest.Cmd = POST_APPEND_OBJECT
	case ossRequest.Query.Has("uploadId"):
		ossRequest.Cmd = POST_COMPLETE_MULTIPART_UPLOAD
	case ossRequest.Query.Has("delete"):
		ossRequest.Cmd = DELETE_MULTIPLE_OBJECTS
	case ossRequest.Query.Has("restore"):
		ossRequest.Cmd = POST_RESTORE_OBJECT
	// TODO 这个有问题
	case ossRequest.Query.Has("multipart/form-data; boundary=(.+)"):
		ossRequest.Cmd = POST_OBJECT
	default:
		ossRequest.Cmd = POST_ELSE
	}
}

func (ossRequest *OssRequest) ParsePut() {
	if ossRequest.Path == "/" {
		ossRequest.Cmd = PUT_BUCKET
		return
	}

	elems := strings.Split(ossRequest.Path[1:], "/")
	ossRequest.Bucket = elems[0]

    if len(elems) < 2 || (len(elems) == 2 && elems[1] == "") {
		switch {
		case ossRequest.Query.Has("acl"):
			ossRequest.Cmd = PUT_BUCKET_ACL
		case ossRequest.Query.Has("logging"):
			ossRequest.Cmd = PUT_BUCKET_LOGGING
		case ossRequest.Query.Has("website"):
			ossRequest.Cmd = PUT_BUCKET_WEBSITE
		case ossRequest.Query.Has("referer"):
			ossRequest.Cmd = PUT_BUCKET_REFERER
		case ossRequest.Query.Has("lifecycle"):
			ossRequest.Cmd = PUT_BUCKET_LIFECYCLE
		// TODO 少一个
		// case ossRequest.Header.Get("x-oss-acl"):
		//     ossRequest.Cmd = PUT_BUCKET_ACL
		default:
			ossRequest.Cmd = PUT_BUCKET
		}
	} else {
		switch {
		case ossRequest.Query.Has("acl"):
			ossRequest.Cmd = PUT_BUCKET_ACL
		case ossRequest.Query.Has("symlink"):
			ossRequest.Cmd = PUT_SYMLINK
		case ossRequest.Query.Has("partNumber="):
			// if ossRequest.Header.Get("x-oss-acl") {
			//     ossRequest.Cmd = PUT_COPY_OBJECT
			// } else {
			//     ossRequest.Cmd = PUT_OBJECT
			// }
			ossRequest.Cmd = PUT_UPLOAD_PART
		default:
			ossRequest.Cmd = PUT_OBJECT
		}

	}
	ossRequest.Object = strings.Join(elems[1:], "/")

	// Also parse x-oss-copy-source-range:bytes=first-last header for multipart copy
	copySource := ossRequest.Header.Get("x-oss-copy-source")
	if copySource != "" {
		srcElems := strings.Split(copySource, "/")
		var root_offset int
		if srcElems[0] == "" {
			root_offset = 1
		} else {
			root_offset = 0
		}
		ossRequest.SrcBucket = srcElems[root_offset]
		ossRequest.SrcBucket = strings.Join(srcElems[1+root_offset:], "/")
		ossRequest.Cmd = PUT_COPY_OBJECT
	}
}

// get_bucket_from_header_host

func (ossRequest *OssRequest) InspectRequest() {
	fmt.Printf("%+v\n", ossRequest)
}

func httpTrace(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        _, err := fmt.Fprintln(os.Stdout, "---------START-HTTP---------")
        if err != nil {
            return
        }

        reqTrace, err := httputil.DumpRequest(r, false)
        if err != nil {
          fmt.Println(err)
        }

        _, err = fmt.Fprint(os.Stdout, string(reqTrace))
        if err != nil {
            return
        }

		f.ServeHTTP(w, r)

  //       http.ReadResponse(w, nil)
  //       respTrace, err := httputil.DumpResponse(w, false)
		// if err != nil {
		// 	return
		// }
        _, err = fmt.Fprintf(os.Stdout, "%#v\n", w)
        if err != nil {
            return
        }

        _, err = fmt.Fprintln(os.Stdout, "---------END-HTTP---------")
        if err != nil {
            return
        }
	}
}
