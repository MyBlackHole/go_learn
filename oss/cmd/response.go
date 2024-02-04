package oss

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type OssResponse struct {
}

type Deleted struct {
	Text string `xml:",chardata"`
	Key  string `xml:"Key"`
}

type DeleteResult struct {
	XMLName xml.Name  `xml:"DeleteResult"`
	Text    string    `xml:",chardata"`
	Deleted []Deleted `xml:"Deleted"`
}

type Error struct {
	XMLName    xml.Name `xml:"Error"`
	Text       string   `xml:",chardata"`
	Code       string   `xml:"Code"`
	Message    string   `xml:"Message"`
	RequestId  string   `xml:"RequestId"`
	HostId     string   `xml:"HostId"`
	BucketName string   `xml:"BucketName"`
}

type Owner struct {
	Text        string `xml:",chardata"`
	DisplayName string `xml:"DisplayName"`
	ID          string `xml:"ID"`
}

type AccessControlList struct {
	Text  string `xml:",chardata"`
	Grant string `xml:"Grant"`
}

type XmlBucket struct {
	Text              string            `xml:",chardata"`
	CreationDate      string            `xml:"CreationDate"`
	ExtranetEndpoint  string            `xml:"ExtranetEndpoint"`
	IntranetEndpoint  string            `xml:"IntranetEndpoint"`
	Location          string            `xml:"Location"`
	Name              string            `xml:"Name"`
	Storage           string            `xml:"Storage"`
	Owner             Owner             `xml:"Owner"`
	AccessControlList AccessControlList `xml:"AccessControlList"`
}

type AccessControlPolicy struct {
	XMLName           xml.Name          `xml:"AccessControlPolicy"`
	Text              string            `xml:",chardata"`
	Owner             Owner             `xml:"Owner"`
	AccessControlList AccessControlList `xml:"AccessControlList"`
}

type BucketInfo struct {
	XMLName xml.Name  `xml:"BucketInfo"`
	Text    string    `xml:",chardata"`
	Bucket  XmlBucket `xml:"Bucket"`
}

func (ossResponse *OssResponse) ResponseError(response http.ResponseWriter, code ErrorCode) {
    requestId := GenerateRequestID()
	response.Header().Set("Server", ALIYUN_OSS_SERVER)
	response.Header().Set("x-oss-request-id", requestId)
	response.WriteHeader(code.StatusCode)
	response.Header().Set("Content-Type", "application/xml")

	switch code.ErrorCode {
	case NO_SUCH_BUCKET.ErrorCode, BUCKET_NOT_EMPTY.ErrorCode:
		errMessage := Error{
			Code:       code.ErrorCode,
			Message:    code.Message,
			RequestId:  requestId,
			HostId:     HOST,
			BucketName: code.Bucket,
		}
		out, err := xml.Marshal(errMessage)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
	default:
		errMessage := Error{
			Code:      code.ErrorCode,
			Message:   code.Message,
			RequestId: requestId,
			HostId:    HOST,
		}
		out, err := xml.Marshal(errMessage)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
	}
}

// Response to get_object by chunk
func (ossResponse *OssResponse) ResponseGetObjectByChunk(response http.ResponseWriter, dataset *Dataset) {
	response.Header().Set("Server", ALIYUN_OSS_SERVER)
	response.Header().Set("x-oss-request-id", GenerateRequestID())
	response.Header().Set("Content-Type", dataset.ContentType)

	if dataset.ContentEncoding != "" {
		response.Header().Set("X-Content-Encoding", dataset.ContentEncoding)
		response.Header().Set("Content-Encoding", dataset.ContentEncoding)
	}

	if dataset.ContentDisposition != "" {
		response.Header().Set("Content-Disposition", dataset.ContentDisposition)
	}

	response.Header().Set("Last-Modified", dataset.ModifiedDate)
	response.Header().Set("ETag", dataset.Md5)
	response.Header().Set("Accept-Ranges", "bytes")
	response.Header().Set("Access-Control-Allow-Origin", "*")

	for k, v := range dataset.CustomMetadata {
		response.Header().Set(fmt.Sprintf("x-oss-meta-%s", k), v)
	}

	objectDir := filepath.Join(GConfig.StoreRoot, dataset.Bucket, dataset.Object)
    objectContentFilename := filepath.Join(objectDir, OBJECT_CONTENT)

	if !dataset.Multipart {
		if dataset.ContentRange != "" {
			response.WriteHeader(206)
			response.Header().Set("Content-Range", dataset.ContentRange)
            file, err := os.Open(objectContentFilename)
            if err != nil {
                panic(err)
            }

            buf := make([]byte, dataset.BytesToRead)

            _, err = file.ReadAt(buf, dataset.Pos)
            if err != nil {
                panic(err)
            }

            response.Write(buf)
		} else {
            file, err := os.Open(objectContentFilename)
            if err != nil {
                panic(err)
            }

            _, err = io.Copy(response, file)
            if err != nil {
                panic(err)
            }
		}
	} else {
        // TODO 暂时不处理这情况
		if dataset.ContentRange != "" {
			response.WriteHeader(206)
			response.Header().Set("Content-Range", dataset.ContentRange)
			// TODO 读取文件
			// 写 body
			// options = { type: 'multipart_range', start_pos: dataset[:pos], read_length: dataset[:bytes_to_read], request_id: request_id,
			//             base_part_filename: File.join(object_dir, "#{Store::OBJECT_CONTENT_PREFIX}")
			//           }
			// response.body = ChunkFile.open(object_content_filename, options)

		} else {

			// TODO 读取文件
			// 写 body
			// options = { type: 'multipart_whole', request_id: request_id, base_part_filename: File.join(object_dir, "#{Store::OBJECT_CONTENT_PREFIX}") }
			// response.body = ChunkFile.open(object_content_filename, options)
		}
	}
}

type InitiateMultipartUploadResult struct {
	XMLName  xml.Name `xml:"InitiateMultipartUploadResult"`
	Text     string   `xml:",chardata"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	UploadId string   `xml:"UploadId"`
}

type CompleteMultipartUploadResult struct {
	XMLName  xml.Name `xml:"CompleteMultipartUploadResult"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Location string   `xml:"Location"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	ETag     string   `xml:"ETag"`
} 

type PostResponse struct {
	XMLName  xml.Name `xml:"PostResponse"`
	Text     string   `xml:",chardata"`
	Location string   `xml:"Location"`
	Bucket   string   `xml:"Bucket"`
	Key      string   `xml:"Key"`
	ETag     string   `xml:"ETag"`
} 

// Response OK to various Request Method
func (ossResponse *OssResponse) ResponseOk(response http.ResponseWriter, dataset *Dataset) {
    // 2024-02-02 15:18 不能直接调用, 不然后面无法写入头了
	// response.WriteHeader(200)
    requestId := GenerateRequestID()
	response.Header().Set("Server", ALIYUN_OSS_SERVER)
	response.Header().Set("x-oss-request-id", requestId)

	switch dataset.Cmd {
	case PUT_BUCKET:
		response.Header().Set("Location", "oss-example")
		response.Header().Set("Access-Control-Allow-Origin", "*")
	case GET_BUCKET_INFO:
		response.Header().Set("x-oss-server-time", time.Now().Format(time.RFC1123))
		response.Header().Set("Content-Type", "application/xml")

		bucketInfo := BucketInfo{
			Bucket: XmlBucket{
				CreationDate:     dataset.CreationDate,
				ExtranetEndpoint: "oss-cn-hangzhou-zmf.aliyuncs.com",
				IntranetEndpoint: "oss-cn-hangzhou-zmf-internal.aliyuncs.com",
				Location:         "cn-hangzhou",
				Name:             dataset.Bucket,
				Storage:          "Standard",
				Owner: Owner{
					DisplayName: "1390402650033793",
					ID:          "1390402650033793",
				},
				AccessControlList: AccessControlList{
					Grant: dataset.ACL,
				},
			},
		}
		out, err := xml.Marshal(bucketInfo)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
	case GET_BUCKET_ACL:
		response.Header().Set("Content-Type", "application/xml")
		accessControlPolicy := AccessControlPolicy{
			Owner: Owner{
				ID:          OWNER_ID,
				DisplayName: OWNER_DISPLAY_NAME,
			},
			AccessControlList: AccessControlList{
				Grant: "private",
			},
		}
		out, err := xml.Marshal(accessControlPolicy)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
	case GET_BUCKET_LOCATION:
        // TODO 暂时不需要支持
		response.Header().Set("Content-Type", "application/xml")
		// <?xml version="1.0" encoding="UTF-8"?>
		// <LocationConstraint>oss-cn-hangzhou</LocationConstraint>
	case PUT_OBJECT:
		response.Header().Set("ETag", "00000000000")
		response.Header().Set("x-oss-bucket-version", "1418321259")
	case POST_APPEND_OBJECT:
		response.Header().Set("x-oss-next-append-position", fmt.Sprintf("%d", dataset.Size))
	case DELETE_OBJECT:
		response.WriteHeader(204)
	case DELETE_MULTIPLE_OBJECTS:
        // TODO 暂时不需要支持
        // dataset.include?(:object_list)

        // xmlDeleledList := ""
        // dataset[:object_list].each do |obj_name|
        //   xml_deleled_list += "<Deleted><Key>#{obj_name}</Key></Deleted>"
        // end
        // output = <<-eos.strip
        //   <?xml version="1.0" encoding="UTF-8"?><DeleteResult>#{xml_deleled_list}</DeleteResult>
        // eos
        // response.body = output
	case HEAD_OBJECT:
		response.Header().Set("Accept-Ranges", "bytes")
		response.Header().Set("ETag", "00000000000")
		response.Header().Set("Content-Length", fmt.Sprintf("%d", dataset.Size))
		// response.Header().Set("Last-Modified", dataset.ModifiedDate)
		response.Header().Set("x-oss-object-type", "Appendable")
		response.Header().Set("x-oss-storage-class", "Standard")
        response.Header().Set("x-oss-server-time", time.Now().Format(time.RFC1123))
		response.Header().Set("x-oss-next-append-position", fmt.Sprintf("%d", dataset.Size))
    case POST_INIT_MULTIPART_UPLOAD:
        // TODO 暂时用不上
		response.Header().Set("Content-Type", "application/xml")
		initiateMultipartUploadResult := InitiateMultipartUploadResult{
            Bucket: dataset.Bucket,
            Key: dataset.Object,
            // TODO 上传 id 
            UploadId: "0000000",
		}
		out, err := xml.Marshal(initiateMultipartUploadResult)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
    case POST_COMPLETE_MULTIPART_UPLOAD:
		response.Header().Set("Content-Type", "application/xml")
		completeMultipartUploadResult := CompleteMultipartUploadResult{
            Location: "Location",
            Bucket: dataset.Bucket,
            Key: dataset.Object,
            // TODO md5
            ETag: "00000000000",
		}
		out, err := xml.Marshal(completeMultipartUploadResult)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
    case GET_OBJECT:
        ossResponse.ResponseGetObjectByChunk(response, dataset)
    case PUT_COPY_OBJECT:
		response.Header().Set("Content-Type", "application/xml")
        // TODO 暂时用不上
        // output = ""
        // xml = Builder::XmlMarkup.new(:target => output)
        // xml.instruct! :xml, :version=>"1.0", :encoding=>"UTF-8"
        // xml.CopyObjectResult(:xmlns => HttpMsg::XMLNS) { |result|
        //   result.LastModified(dataset[:modified_date])
        //   result.ETag("\"#{dataset[:md5]}\"")
        // }
        // response.body = output
    case GET_OBJECT_META:
        // TODO 根据 api 文档处理的
		response.Header().Set("Content-Length", fmt.Sprintf("%d", dataset.Size))
		response.Header().Set("Last-Modified", dataset.ModifiedDate)
		response.Header().Set("ETag", "00000000000")
    case GET_OBJECT_ACL:
		response.Header().Set("Content-Type", "application/xml")
		response.Header().Set("x-oss-server-time", dataset.ModifiedDate)
		accessControlPolicy := AccessControlPolicy{
			Owner: Owner{
				ID:          OWNER_ID,
				DisplayName: OWNER_DISPLAY_NAME,
			},
			AccessControlList: AccessControlList{
				Grant: "default",
			},
		}
		out, err := xml.Marshal(accessControlPolicy)
		if err != nil {
			panic(err)
		}
		_, err = response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}

		_, err = response.Write(out)
		if err != nil {
			panic(err)
		}
    case POST_OBJECT:
        // TODO 根据 api 文档不需要处理
        // 支持表单文件上传
		response.Header().Set("ETag", "00000000000")
    case GET_BUCKET:
		response.Header().Set("Content-Type", "application/xml")
        _, err := response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}
        response.Write(dataset.Body)
    case LIST_BUCKETS:
		response.Header().Set("Content-Type", "application/xml")
        _, err := response.Write([]byte(xml.Header))
		if err != nil {
			panic(err)
		}
        response.Write(dataset.Body)
    default:
        if dataset.Body != nil && len(dataset.Body) > 0 {
            response.Write(dataset.Body)
        }
	}
}

// Response while request method is OPTIONS 
func (ossResponse *OssResponse) ResponseOptions(response http.ResponseWriter) {
    response.Header().Set("Access-Control-Allow-Origin", "*")
    response.Header().Set("Access-Control-Allow-Methods", "PUT, POST, HEAD, GET, OPTIONS")
    response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization, Content-Length, ETag, X-CSRF-Token, Content-Disposition")
    response.Header().Set("Access-Control-Expose-Headers", "ETag")
}


// NoSuchBucket
func (ossResponse *OssResponse) ResponseNoSuchBucket(response http.ResponseWriter, bucket string) bool {
    isExist, err := pathExists(filepath.Join(GConfig.StoreRoot, bucket, BUCKET_METADATA))
	if err != nil {
		panic(err)
	}

    if !isExist {
        code := NO_SUCH_BUCKET
        code.Bucket = bucket
        ossResponse.ResponseError(response, code)
        return true
    }
    return false
    
}

// NoSuchBucket when DeleteBucket
func (ossResponse *OssResponse) ResponseNoSuchBucketWhenDeleteBucket(response http.ResponseWriter, bucket string) bool {
    bucketDir := filepath.Join(GConfig.StoreRoot, bucket)
    isExist, err := pathExists(filepath.Join(bucketDir, BUCKET_METADATA))
	if err != nil {
		panic(err)
	}

    if !isExist {
        os.RemoveAll(bucketDir)

        code := NO_SUCH_BUCKET
        code.Bucket = bucket
        ossResponse.ResponseError(response, code)

        return true
    }
    return false
}

// BucketNotEmpty
func (ossResponse *OssResponse) ResponseBucketNotEmpty(response http.ResponseWriter, bucket string) bool {
    bucketDir := filepath.Join(GConfig.StoreRoot, bucket)
    dirs, files, err := listDir(bucketDir)
    if err != nil {
        panic(err)
    }
    if len(dirs) > 0 || len(files) > 0 {
        code := BUCKET_NOT_EMPTY
        code.Bucket = bucket
        ossResponse.ResponseError(response, code)
        return true
    }
    return false
}

// InvalidBucketName
func (ossResponse *OssResponse) ResponseInvalidBucketName(response http.ResponseWriter, bucket string) bool {
    if !ValidBucketName(bucket) {
        code := INVALID_BUCKET_NAME
        code.Bucket = bucket
        ossResponse.ResponseError(response, code)
        return true
    }
    return false
}

// TooManyBuckets
func (ossResponse *OssResponse) ResponseTooManyBuckets(response http.ResponseWriter) bool {
    var bucketCount int
    dirs, _, err := listDir(GConfig.StoreRoot)
    if err != nil {
        panic(err)
    }

    for _, dir := range dirs {
        isExist, err := pathExists(filepath.Join(dir, BUCKET_METADATA))
		if err != nil {
			panic(err)
		}
		if isExist {
			bucketCount += 1
		}
    }

    if bucketCount >= MAX_BUCKET_NUM {
        ossResponse.ResponseError(response, TOO_MANY_BUCKETS)
        return true
    }

    return false
}

// InvalidObjectName
func (ossResponse *OssResponse) ResponseInvalidObjectName(response http.ResponseWriter, object string) bool {
    if !ValidObjectName(object) {
        code := INVALID_OBJECT_NAME
        code.Object = object
        ossResponse.ResponseError(response, code)
        return true
    }
    return false
}

// NoSuchKey/NotFound
func (ossResponse *OssResponse) ResponseNoSuchObject(response http.ResponseWriter, bucket, object string) bool {
    isExist, err := pathExists(filepath.Join(GConfig.StoreRoot, bucket, object, OBJECT_METADATA))
    if err != nil {
        panic(err)
    }
    if !isExist {
        code := NOT_FOUND
        code.Bucket = bucket
        code.Object = object
        ossResponse.ResponseError(response, code)
        return true
    }
    return false
}
