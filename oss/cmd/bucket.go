package oss

import (
	"encoding/xml"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Bucket             string            `yaml:":bucket"`
	Object             string            `yaml:":object"`
	ACL                string            `yaml:":acl"`
	CreationDate       string            `yaml:":creation_date"`
	ModifiedDate       string            `yaml:":modified_date"`
	ContentType        string            `yaml:":content_type"`
	ContentDisposition string            `yaml:":content_disposition"`
	ContentEncoding    string            `yaml:":content_encoding"`
	Size               int               `yaml:":size"`
	PartSize           int               `yaml:":part_size"`
	Md5                string            `yaml:":md5"`
	Appendable         bool              `yaml:":appendable"`
	OssMetadata        map[string]string `yaml:":oss_metadata"`
	CustomMetadata     map[string]string `yaml:":custom_metadata"`
}

type Dataset struct {
	Cmd                string
	Bucket             string
	Object             string
	Md5                string
	Multipart          bool
	ContentType        string
	ContentLength      int
	ContentRange       string
	ContentDisposition string
	ContentEncoding    string
	Size               int
	PartSize           int
	CreationDate       string
	ModifiedDate       string

	Pos         int64
	BytesToRead int

	ACL string

	Appendable bool

	OssMetadata    map[string]string
	CustomMetadata map[string]string
	Body           []byte
}

type Bucket struct {
	Name         string
	CreationDate string
}

type ListBucketsResponse struct {
	XMLName xml.Name `xml:"http://s3.amazonaws.com/doc/2006-03-01/ ListAllMyBucketsResult" json:"-"`
	Xmlns   string   `xml:"xmlns,attr"`

	Owner Owner

	Buckets struct {
		Buckets []Bucket `xml:"Bucket"`
	}
}

type OssBucket struct {
}

// GetService=ListBuckets
func (ossbucket *OssBucket) GetService(response http.ResponseWriter) {
	ossbucket.ListBuckets(response)
}

// ListBuckets=GetService
func (ossbucket *OssBucket) ListBuckets(response http.ResponseWriter) {
	ossResponse := OssResponse{}
	data := ListBucketsResponse{}
	data.Xmlns = XMLNS
	data.Owner = Owner{
		ID:          "00220120222",
		DisplayName: "1390402650033793",
	}
	buckets, _, err := listDir(GConfig.StoreRoot)
	if err != nil {
		panic(err)
	}

	for _, bucket := range buckets {
		bucketDir := filepath.Join(GConfig.StoreRoot, bucket)
		bucketMetadataFile := filepath.Join(bucketDir, BUCKET_METADATA)
		isExist, err := pathExists(bucketMetadataFile)
		if err != nil {
			panic(err)
		}
		if !isExist {
			continue
		}
		cTimespec := GetFileCreateTime(bucketDir)
		cTim := TimespecToTime(cTimespec)
		bucketNode := Bucket{
			Name:         bucket,
			CreationDate: cTim.UTC().Format(time.RFC3339),
		}
		data.Buckets.Buckets = append(data.Buckets.Buckets, bucketNode)
	}

	out, err := xml.Marshal(data)
	if err != nil {
		panic(err)
	}

	dataset := Dataset{}
	dataset.Cmd = LIST_BUCKETS
	dataset.Body = out
	ossResponse.ResponseOk(response, &dataset)
}

// PutBucket=CreateBucket
func (ossbucket *OssBucket) CreateBucket(bucket string, request *http.Request, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseInvalidBucketName(response, bucket) {
		return
	}

	if ossResponse.ResponseTooManyBuckets(response) {
		return
	}

	bucketFoler := filepath.Join(GConfig.StoreRoot, bucket)
	bucketMetadataFile := filepath.Join(bucketFoler, BUCKET_METADATA)
	isExist1, err := pathExists(bucketFoler)
	if err != nil {
		panic(err)
	}

	isExist2, err := pathExists(bucketMetadataFile)
	if err != nil {
		panic(err)
	}

	if !isExist1 && !isExist2 {
		// 创建多级目录
		os.MkdirAll(bucketFoler, os.ModePerm)
		cTimespec := GetFileCreateTime(bucketFoler)
		cTim := TimespecToTime(cTimespec)
		metadata := Metadata{
			Bucket:       bucket,
			CreationDate: cTim.UTC().Format(time.RFC3339),
			ACL:          request.Header.Get("x-oss-acl"),
		}
		out, err := yaml.Marshal(metadata)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(bucketMetadataFile, out, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	dataset := Dataset{}
	dataset.Cmd = PUT_BUCKET
	ossResponse.ResponseOk(response, &dataset)
}

// PutBucketACL
func (ossbucket *OssBucket) PutBucketAcl(response http.ResponseWriter) {
	ossResponse := OssResponse{}
	dataset := Dataset{}
	ossResponse.ResponseOk(response, &dataset)
}

type Content struct {
	Text         string `xml:",chardata"`
	Key          string `xml:"Key"`
	// LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Type         string `xml:"Type"`
	Size         int `xml:"Size"`
	StorageClass string `xml:"StorageClass"`
	Owner        Owner  `xml:"Owner"`
}

// ListBucketResult was generated 2024-01-30 11:26:06 by https://xml-to-go.github.io/ in Ukraine.
type ListBucketResult struct {
	XMLName      xml.Name  `xml:"ListBucketResult"`
	Text         string    `xml:",chardata"`
	Xmlns        string    `xml:"xmlns,attr"`
	Name         string    `xml:"Name"`
	Prefix       string    `xml:"Prefix"`
	Marker       string    `xml:"Marker"`
	MaxKeys      string    `xml:"MaxKeys"`
	Delimiter    string    `xml:"Delimiter"`
	EncodingType string    `xml:"EncodingType"`
	NextMarker   string    `xml:"NextMarker"`
	IsTruncated  bool    `xml:"IsTruncated"`
	Contents     []Content `xml:"Contents"`
}

// GetBucket=ListObjects
func (ossbucket *OssBucket) GetBucket(bucket string, ossRequest *OssRequest, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchBucket(response, bucket) {
		return
	}

	vars := ossRequest.Request.URL.Query()
	marker := vars.Get("marker")
	prefix := vars.Get("prefix")
	maxKeys := vars.Get("max-keys")
	delimiter := vars.Get("delimiter")

	data := &ListBucketResult{}
	data.Xmlns = XMLNS
	data.Name = bucket
	data.Marker = marker
	data.Prefix = prefix
	data.MaxKeys = maxKeys
	data.Delimiter = delimiter
	data.EncodingType = "url"

	GetBucketListObjects(data, ossRequest)

    out, err := xml.Marshal(data)
    if err != nil {
        panic(err)
    }
	dataset := Dataset{}
	dataset.Cmd = GET_BUCKET
    dataset.Body = out
	ossResponse.ResponseOk(response, &dataset)
}

// GetBucketACL
func (ossbucket *OssBucket) GetBucketAcl(bucket string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	if ossResponse.ResponseNoSuchBucket(response, bucket) {
		return
	}

	dataset := Dataset{}
	dataset.Cmd = GET_BUCKET_ACL
	ossResponse.ResponseOk(response, &dataset)
}

// GetBucketLocation:
func (ossbucket *OssBucket) GetBucketLocation(response http.ResponseWriter) {
	ossResponse := OssResponse{}
	dataset := Dataset{}
	dataset.Cmd = GET_BUCKET_LOCATION
	ossResponse.ResponseOk(response, &dataset)
}

// GetBucketInfo
func (ossbucket *OssBucket) GetBucketInfo(bucket string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	// NoSuchBucket
	if ossResponse.ResponseNoSuchBucket(response, bucket) {
		return
	}

	bucketMetadataFile := filepath.Join(GConfig.StoreRoot, bucket, BUCKET_METADATA)
	metadata := Metadata{}
	ret, err := os.ReadFile(bucketMetadataFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(ret, &metadata)
	if err != nil {
		panic(err)
	}
	dataset := Dataset{}
	dataset.Cmd = GET_BUCKET_INFO
	dataset.Bucket = bucket
	dataset.ACL = "private"
	ossResponse.ResponseOk(response, &dataset)
}

// DeleteBucket
func (ossbucket *OssBucket) DeleteBucket(bucket string, response http.ResponseWriter) {
	ossResponse := OssResponse{}
	// NoSuchBucket
	if ossResponse.ResponseNoSuchBucketWhenDeleteBucket(response, bucket) {
		return
	}

	// BucketNotEmpty
	if ossResponse.ResponseBucketNotEmpty(response, bucket) {
		return
	}

	// DeleteBucketFolder
	bucketDir := filepath.Join(GConfig.StoreRoot, bucket)
	err := os.RemoveAll(bucketDir)
	if err != nil {
		panic(err)
	}

	dataset := Dataset{}
	ossResponse.ResponseOk(response, &dataset)
}
