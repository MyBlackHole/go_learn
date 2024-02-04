package oss

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type OssObject struct {
}

// PutObject
func (ossObject *OssObject) PutObject(bucket, object string, request *http.Request, response http.ResponseWriter, partNumber int) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchBucket(response, bucket) {
		return
	}

	if ossResponse.ResponseInvalidObjectName(response, object) {
		return
	}

	//    // TODO 不需要处理
	// var checkChunkedFilesize bool
	var objectContentFilename string

	contentLength := request.ContentLength
	transferEncoding := request.TransferEncoding

	if contentLength > -1 {
		if contentLength > int64(MAX_OBJECT_FILE_SIZE) {
			ossResponse.ResponseError(response, INVALID_ARGUMENT)
			return
		}
	} else {

		if len(transferEncoding) > 0 {
			if transferEncoding[0] != "chunked" {
				ossResponse.ResponseError(response, MISSING_CONTENT_LENGTH)
				return
			} else {
				// checkChunkedFilesize = true
			}
		} else {
			ossResponse.ResponseError(response, MISSING_CONTENT_LENGTH)
			return
		}

	}

	objDir := filepath.Join(GConfig.StoreRoot, bucket, object)
	if partNumber > 0 {
		objectContentFilename = filepath.Join(objDir, fmt.Sprintf("%s%s", OBJECT_CONTENT_PREFIX, partNumber))
	} else {
		DeleteObjectFileAndDir(bucket, object)
		objectContentFilename = filepath.Join(objDir, OBJECT_CONTENT)
	}

	err := os.MkdirAll(objDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fileObjectContent, err := os.OpenFile(objectContentFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// TODO 先不考虑这情况
	// contentType := request.Header.Get("Content-Type")
	// request.MultipartReader()
	_, err = io.Copy(fileObjectContent, request.Body)
	if err != nil {
		panic(err)
	}

	err = fileObjectContent.Sync()
	if err != nil {
		panic(err)
	}

	err = fileObjectContent.Close()
	if err != nil {
		panic(err)
	}

	PutObjectMetadata(bucket, object, request, nil)
	dataset := Dataset{}
	dataset.Cmd = PUT_OBJECT
	ossResponse.ResponseOk(response, &dataset)
}

// CopyObject
func (ossObject *OssObject) CopyObject(srcBucket, srcObject string, dstBucket, dstObject string, request *http.Request, response http.ResponseWriter) {
	srcObjectDir := filepath.Join(GConfig.StoreRoot, srcBucket, srcObject)
	srcMetadataFilename := filepath.Join(srcObjectDir, OBJECT_METADATA)

	dstObjectDir := filepath.Join(GConfig.StoreRoot, dstBucket, dstObject)
	dstMetadataFilename := filepath.Join(dstObjectDir, OBJECT_METADATA)

	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchBucket(response, srcBucket) {
		return
	}
	if ossResponse.ResponseNoSuchObject(response, srcBucket, srcObject) {
		return
	}

	if srcBucket == dstBucket && srcObject == srcObject {
		PutObjectMetadata(dstBucket, dstObject, request, nil)
		// TODO 暂时用不上
		dataset := Dataset{}
		dataset.Cmd = PUT_COPY_OBJECT
		ossResponse.ResponseOk(response, &dataset)
		return
	}

	dstBucketMetadataFile := filepath.Join(GConfig.StoreRoot, dstBucket, BUCKET_METADATA)
	exist, err := pathExists(dstBucketMetadataFile)
	if err != nil {
		panic(err)
	}
	if !exist {
		dstBucketDir := filepath.Join(GConfig.StoreRoot, dstBucket)
		err = os.MkdirAll(dstBucketDir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		out, err := os.ReadFile(srcMetadataFilename)
		metadata := Metadata{}
		err = yaml.Unmarshal(out, &metadata)
		if err != nil {
			panic(err)
		}

		metadata.Bucket = dstBucket
		metadata.CreationDate = time.Now().Format(time.RFC1123)

		out, err = yaml.Marshal(metadata)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(dstBucketMetadataFile, out, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	exist, err = pathExists(dstBucketMetadataFile)

	if err != nil {
		panic(err)
	}

	if !exist {
		err = os.MkdirAll(dstObjectDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	fileNumber := 1
	srcContentFilenameBase := filepath.Join(srcObjectDir, OBJECT_CONTENT_PREFIX)
	dstContentFilenameBase := filepath.Join(dstObjectDir, OBJECT_CONTENT_PREFIX)

	for {
		currentSrcFilename := fmt.Sprintf("%s%d", srcContentFilenameBase, fileNumber)
		exist, err = pathExists(currentSrcFilename)

		if err != nil {
			panic(err)
		}

		if !exist {
			// 遍历完成退出
			break
		}

		currentDstFilename := fmt.Sprintf("%s%d", dstContentFilenameBase, fileNumber)
		srcFile, err := os.OpenFile(currentSrcFilename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(currentDstFilename, os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			panic(err)
		}
		fileNumber += 1
	}

	var metadata *Metadata
	metadataDirective := request.Header.Get("x-oss-metadata-directive")
	if metadataDirective == "REPLACE" {
		metadata = PutObjectMetadata(dstBucket, dstObject, request, nil)
	} else {
		srcFile, err := os.OpenFile(srcMetadataFilename, os.O_RDONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(dstMetadataFilename, os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			panic(err)
		}
		srcFile.Seek(0, io.SeekStart)

		out, err := io.ReadAll(srcFile)
		metadata := Metadata{}
		err = yaml.Unmarshal(out, &metadata)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("%+v", metadata)

	// TODO 暂时用不上
	dataset := Dataset{}
	dataset.Cmd = PUT_COPY_OBJECT
	ossResponse.ResponseOk(response, &dataset)
}

// GetObject
func (ossObject *OssObject) GetObject(ossRequest *OssRequest, request *http.Request, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchObject(response, ossRequest.Bucket, ossRequest.Object) {
		return
	}

	objectMultipartContentTag := filepath.Join(GConfig.StoreRoot, ossRequest.Bucket, ossRequest.Object, OBJECT_CONTENT_TWO)
	objectMetadataFilename := filepath.Join(GConfig.StoreRoot, ossRequest.Bucket, ossRequest.Object, OBJECT_METADATA)

	out, err := os.ReadFile(objectMetadataFilename)
	metadata := Metadata{}
	err = yaml.Unmarshal(out, &metadata)
	if err != nil {
		panic(err)
	}

	dataset := Dataset{}
	dataset.Cmd = GET_OBJECT
	dataset.Bucket = ossRequest.Bucket
	dataset.Object = ossRequest.Object
	dataset.Md5 = metadata.Md5

	exist, err := pathExists(objectMultipartContentTag)
	if err != nil {
		panic(err)
	}
	if exist {
		dataset.Multipart = true
	} else {
		dataset.Multipart = false
	}

	if metadata.ContentType != "" {
		dataset.ContentType = "application/octet-stream"
	}

	dataset.ContentDisposition = metadata.ContentDisposition
	dataset.ContentEncoding = metadata.ContentEncoding
	dataset.Size = metadata.Size
	dataset.PartSize = metadata.PartSize
	dataset.CreationDate = metadata.CreationDate
	dataset.ModifiedDate = metadata.ModifiedDate
	dataset.CustomMetadata = metadata.CustomMetadata
	contentLength := dataset.Size
	rangeString := request.Header.Get("range")
	if rangeString != "" {
		if strings.HasPrefix(rangeString, "bytes=") {
			byteRangeString := strings.TrimPrefix(rangeString, "bytes=")
			rets := strings.Split(byteRangeString, "-")
			if len(rets) != 2 {
				panic("rets not 2")
			}
			start, err := strconv.Atoi(rets[0])
			if err != nil {
				panic(err)
			}
			finish, err := strconv.Atoi(rets[1])
			if err != nil {
				panic(err)
			}

			if finish == 0 {
				finish = contentLength - 1
			}
			dataset.Pos = int64(start)
			dataset.BytesToRead = finish - start + 1
			dataset.ContentRange = fmt.Sprintf("bytes %d-%d/%d", start, finish, contentLength)
		}

	} else {
		dataset.ContentLength = dataset.Size
	}
	ossResponse.ResponseGetObjectByChunk(response, &dataset)
}

// AppendObject
func (ossObject *OssObject) AppendObject(bucket, object string, request *http.Request, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchBucket(response, bucket) {
		return
	}

	objDir := filepath.Join(GConfig.StoreRoot, bucket, object)
	metadataFilename := filepath.Join(objDir, OBJECT_METADATA)

	exist, err := pathExists(metadataFilename)
	if err != nil {
		panic(err)
	}

	if exist {
		out, err := os.ReadFile(metadataFilename)
		metadata := Metadata{}
		err = yaml.Unmarshal(out, &metadata)
		if err != nil {
			panic(err)
		}

		if !metadata.Appendable {
			ossResponse.ResponseError(response, OBJECT_NOT_APPENDABLE)
			return
		}
	}

	err = os.MkdirAll(objDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	contentFilename := filepath.Join(objDir, OBJECT_CONTENT)
	fileObjectContent, err := os.OpenFile(contentFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileObjectContent, request.Body)
	if err != nil {
		panic(err)
	}

	err = fileObjectContent.Sync()
	if err != nil {
		panic(err)
	}

	err = fileObjectContent.Close()
	if err != nil {
		panic(err)
	}

	dataset := Dataset{}
	dataset.Appendable = true
	dataset.Cmd = POST_APPEND_OBJECT
	metadata := PutObjectMetadata(bucket, object, request, &dataset)
	dataset.Size = metadata.Size
	ossResponse.ResponseOk(response, &dataset)
}

// DeleteObject
func (ossObject *OssObject) DeleteObject(bucket, object string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchBucket(response, bucket) {
		return
	}

	DeleteObjectFileAndDir(bucket, object)
	dataset := Dataset{}
	dataset.Cmd = DELETE_OBJECT
	ossResponse.ResponseOk(response, &dataset)
}

// DeleteMultipleObjects
func (ossObject *OssObject) DeleteMultipleObject(bucket string, request *http.Request, response http.ResponseWriter) {
	// TODO 暂不处理
	// xml = Document.new(request.body)
	// quiet = xml.root.elements['/Delete/Quiet'].text
	// object_list = []
	// xml.elements.each('*/Object/Key') do |e|
	//   object = e.text
	//   object_list << object
	//   Object.delete_object(bucket, object, response)
	// end
	//
	// dataset = { cmd: Request::DELETE_MULTIPLE_OBJECTS }
	// dataset[:object_list] = object_list if quiet.downcase == 'false'
	//
	// OssResponse.response_ok(response, dataset)
}

// HeadObject
func (ossObject *OssObject) HeadObject(bucket, object string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchObject(response, bucket, object) {
		return
	}
	objectMetadataFilename := filepath.Join(GConfig.StoreRoot, bucket, object, OBJECT_METADATA)
	out, err := os.ReadFile(objectMetadataFilename)
	metadata := Metadata{}
	err = yaml.Unmarshal(out, &metadata)
	if err != nil {
		panic(err)
	}

	dataset := Dataset{}
	dataset.Cmd = HEAD_OBJECT
	dataset.Size = metadata.Size
	ossResponse.ResponseOk(response, &dataset)
}

// GetObjectMeta
func (ossObject *OssObject) GetObjectMeta(bucket, object string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchObject(response, bucket, object) {
		return
	}
	objectMetadataFilename := filepath.Join(GConfig.StoreRoot, bucket, object, OBJECT_METADATA)
	out, err := os.ReadFile(objectMetadataFilename)
	metadata := Metadata{}
	err = yaml.Unmarshal(out, &metadata)
	if err != nil {
		panic(err)
	}
	dataset := Dataset{}
	dataset.Cmd = GET_OBJECT_META
	ossResponse.ResponseOk(response, &dataset)
}

func (ossObject *OssObject) PutObjectAcl(bucket, object string, request *http.Request, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchObject(response, bucket, object) {
		return
	}
	objectMetadataFilename := filepath.Join(GConfig.StoreRoot, bucket, object, OBJECT_METADATA)
	out, err := os.ReadFile(objectMetadataFilename)
	metadata := Metadata{}
	err = yaml.Unmarshal(out, &metadata)
	if err != nil {
		panic(err)
	}
	// TODO 或许不需要管
	// acl_old = dataset[:acl]
	// dataset[:acl] = request.header['x-oss-object-acl'].first || acl_old
	// File.open(object_metadata_filename, 'w') do |f|
	//   f << YAML.dump(dataset)
	// end
	// dataset[:cmd] = Request::PUT_OBJECT_ACL
	//
	// OssResponse.response_ok(response, dataset)
}

func (ossObject *OssObject) GetObjectAcl(bucket, object string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchObject(response, bucket, object) {
		return
	}
	objectMetadataFilename := filepath.Join(GConfig.StoreRoot, bucket, object, OBJECT_METADATA)
	out, err := os.ReadFile(objectMetadataFilename)
	metadata := Metadata{}
	err = yaml.Unmarshal(out, &metadata)
	if err != nil {
		panic(err)
	}
	// TODO 或许不需要管
	// dataset[:cmd] = Request::GET_OBJECT_ACL
	// dataset[:acl] = 'default'
	//
	// OssResponse.response_ok(response, dataset)
}
