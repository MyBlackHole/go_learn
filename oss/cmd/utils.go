package oss

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

func listDir(path string) (dirs []string, files []string, err error) {
	lst, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, val := range lst {
		if val.IsDir() {
			dirs = append(dirs, val.Name())
		} else {
			files = append(files, val.Name())
		}
	}
	return
}

// 判断所给路径文件/文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	//isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

func GetFileCreateTime(path string) syscall.Timespec {
	fileInfo, _ := os.Stat(path)
	stat_t := fileInfo.Sys().(*syscall.Stat_t)
	return stat_t.Ctim
}

func TimespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func DeleteObjectFileAndDir(bucket, object string) {
	bucketDir := filepath.Join(GConfig.StoreRoot, bucket)
	objectDir := filepath.Join(bucketDir, object)
	objectMetadataFilename := filepath.Join(objectDir, OBJECT_METADATA)
	exist, err := pathExists(objectMetadataFilename)
	if err != nil {
		panic(err)
	}

	if exist {
		err := os.RemoveAll(objectDir)
		if err != nil {
			panic(err)
		}
	}
}

func PutObjectMetadata(bucket, object string, request *http.Request, dataset *Dataset) *Metadata {
	objDir := filepath.Join(GConfig.StoreRoot, bucket, object)
	contentFilename := filepath.Join(objDir, OBJECT_CONTENT)
	// TODO 判断是否存在上下文
	exist, err := pathExists(contentFilename)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic("centent 文件不存在")
	}

	metadata := &Metadata{
		Bucket:             bucket,
		Object:             object,
		ACL:                "default",
		CreationDate:       time.Now().Format(time.RFC1123),
		ModifiedDate:       time.Now().Format(time.RFC1123),
		ContentType:        request.Header.Get("Content-Type"),
		ContentDisposition: request.Header.Get("content-disposition"),
		ContentEncoding:    request.Header.Get("content-encoding"),
	}

	fi, err := os.Stat(contentFilename)
	if err != nil {
		panic(err)
	}

	if dataset != nil {
		metadata.Size = int(fi.Size())
		// TODO 暂时不计算
		metadata.Md5 = "00000000000000000000000000000000"
		metadata.Appendable = dataset.Appendable
		// TODO
		// metadata.PartSize = options["part_size"]
	} else {
		metadata.Size = int(fi.Size())
		// TODO 暂时不计算
		metadata.Md5 = "00000000000000000000000000000000"
		metadata.PartSize = 0
	}
	// TODO 忽略 meta 数据
	// request.header.each do |key, value|
	//   match = /^x-oss-([^-]+)-(.*)$/.match(key)
	//   next unless match
	//
	//   if match[1].eql?('meta') && (match_key = match[2])
	//     metadata[:custom_metadata][match_key] = value.join(', ')
	//     next
	//   end
	//   metadata[:oss_metadata][key.gsub(/^x-oss-/, '')] = value.join(', ')
	// end

	metadataFile := filepath.Join(objDir, OBJECT_METADATA)
	out, err := yaml.Marshal(metadata)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(metadataFile, out, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return metadata
}

func GenerateRequestID() string {
	return uuid.New().String()
}

func ValidBucketName(bucket string) bool {
	// TODO 暂不处理
	return true
}

func ValidObjectName(object string) bool {
	// TODO 暂不处理
	return true
}

var TruncatedERR = errors.New("truncated")

func GetBucketListObjects(result *ListBucketResult, ossRequest *OssRequest) {
	var err error
	var markerFound bool
	var maxKeys int
	var count int
	var keyName string
	marker := result.Marker
	if marker == "" {
		markerFound = true
	}

	prefix := result.Prefix

	if result.MaxKeys == "" {
		maxKeys = 100
	} else {
		maxKeys, err = strconv.Atoi(result.MaxKeys)
		if err != nil {
			panic(err)
		}

		if maxKeys > 1000 {
			maxKeys = 1000
		}
	}

	// TODO 暂时不考虑
	// delimiter := result.Delimiter

	bucketPath := filepath.Join(GConfig.StoreRoot, ossRequest.Bucket)
    bucketPath += "/"
	findRootFolder := bucketPath

	err = filepath.Walk(findRootFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			filename := filepath.Base(path)
			if filename != OBJECT_METADATA {
				return nil
			}

			filedir := filepath.Dir(path)
			keyName = strings.Replace(filedir, bucketPath, "", -1)

			if markerFound {
                if prefix == "" || strings.Index(keyName, prefix) == 1 || strings.Index(keyName, prefix) == 0 {
                    count += 1
                    if count <= maxKeys {
                        metadata := Metadata{}
                        ret, err := os.ReadFile(path)
                        if err != nil {
                            panic(err)
                        }
                        err = yaml.Unmarshal(ret, &metadata)
                        if err != nil {
                            panic(err)
                        }

                        content := Content{}
                        content.Key = keyName
                        content.ETag = metadata.Md5
                        content.Size = metadata.Size
                        // content.LastModified = metadata.ModifiedDate
                        content.StorageClass = "Standard"
                        content.Type = "Multipart"
                        content.Owner.ID = "00220120222"
                        content.Owner.DisplayName = "1390402650033793"
                        result.Contents = append(result.Contents, content)
                    } else {
                        result.IsTruncated = true
                        return TruncatedERR
                    }

                }
			}


			if strings.Compare(marker, keyName) <= 0 {
				markerFound = true
			}

			return nil
		},
	)
	result.NextMarker = keyName

	if err != nil {
		switch err {
		case TruncatedERR:
			return
		default:
			panic(err)
		}
	}
}
