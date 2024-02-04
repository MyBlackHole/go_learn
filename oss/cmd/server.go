package oss

import (
	"net/http"
)

type OssServer struct {

}

func (ossServer *OssServer) HeadHandler(w http.ResponseWriter, r *http.Request)  {
    ossRequest := NewOssRequest(r)
    ossRequest.ParseGet()

    if ossRequest.Cmd == "" {
        panic("Cmd 空")
    }
    ossObject := OssObject{}
    ossResponse := OssResponse{}
    switch ossRequest.Cmd {
    case HEAD_OBJECT:
        ossObject.HeadObject(ossRequest.Bucket, ossRequest.Object, w)
    default:
         ossResponse.ResponseError(w, BAD_REQUEST)
    }
}

func (ossServer *OssServer) GetHandler(w http.ResponseWriter, r *http.Request)  {
    ossRequest := NewOssRequest(r)
    ossRequest.ParseGet()

    ossObject := OssObject{}
    ossBucket := OssBucket{}
    ossResponse := OssResponse{}
    switch ossRequest.Cmd {
    case LIST_BUCKETS:
        ossBucket.ListBuckets(w)
    case GET_BUCKET:
        ossBucket.GetBucket(ossRequest.Bucket, ossRequest, w)
    case GET_BUCKET_ACL:
        ossBucket.GetBucketAcl(ossRequest.Bucket, w)
    case GET_BUCKET_INFO:
        ossBucket.GetBucketInfo(ossRequest.Bucket, w)
    case GET_BUCKET_LOCATION:
        ossBucket.GetBucketLocation(w)
    case GET_BUCKET_LOGGING:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case GET_BUCKET_REFERER:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case GET_BUCKET_WEBSITE:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case GET_BUCKET_LIFECYCLE:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case GET_OBJECT_ACL:
        ossObject.GetObjectAcl(ossRequest.Bucket, ossRequest.Object, w)
    case GET_OBJECT_META:
        ossObject.GetObjectMeta(ossRequest.Bucket, ossRequest.Object, w)
    case GET_SYMLINK:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case GET_OBJECT:
        ossObject.GetObject(ossRequest, r, w)
    case GET_LIST_MULTIPART_UPLOADS:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case GET_LIST_PARTS:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    default:
         ossResponse.ResponseError(w, BAD_REQUEST)
    }
}

func (ossServer *OssServer) PutHandler(w http.ResponseWriter, r *http.Request)  {
    ossRequest := NewOssRequest(r)
    ossRequest.ParsePut()

    if ossRequest.Cmd == "" {
        panic("Cmd 空")
    }

    ossObject := OssObject{}
    ossBucket := OssBucket{}
    ossResponse := OssResponse{}
    switch ossRequest.Cmd {
    case PUT_BUCKET:
        ossBucket.CreateBucket(ossRequest.Bucket, r, w)
    case PUT_BUCKET_ACL:
        ossBucket.PutBucketAcl(w)
    case PUT_BUCKET_LOGGING:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case PUT_BUCKET_WEBSITE:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case PUT_BUCKET_REFERER:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case PUT_BUCKET_LIFECYCLE:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case PUT_OBJECT:
        ossObject.PutObject(ossRequest.Bucket, ossRequest.Object, r, w, 0)
    case PUT_OBJECT_ACL:
        ossObject.PutObjectAcl(ossRequest.Bucket, ossRequest.Object, r, w)
    case PUT_SYMLINK:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case PUT_UPLOAD_PART:
        // TODO 暂不需要
    case PUT_UPLOAD_PART_COPY:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case PUT_COPY_OBJECT:
        ossObject.CopyObject(ossRequest.SrcBucket, ossRequest.SrcObject, ossRequest.Bucket, ossRequest.Object, r, w)
    default:
         ossResponse.ResponseError(w, BAD_REQUEST)
    }
}


func (ossServer *OssServer) PostHandler(w http.ResponseWriter, r *http.Request)  {
    ossRequest := NewOssRequest(r)
    ossRequest.ParsePost()

    if ossRequest.Cmd == "" {
        panic("Cmd 空")
    }

    ossObject := OssObject{}
    ossResponse := OssResponse{}
    switch ossRequest.Cmd {
    case POST_INIT_MULTIPART_UPLOAD:
        // TODO 暂不需要
    case POST_APPEND_OBJECT:
        ossObject.AppendObject(ossRequest.Bucket, ossRequest.Object, r, w)
    case POST_COMPLETE_MULTIPART_UPLOAD:
        // TODO 暂不需要
    case DELETE_MULTIPLE_OBJECTS:
        ossObject.DeleteMultipleObject(ossRequest.Bucket, r, w)
    case POST_OBJECT:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    default:
         ossResponse.ResponseError(w, BAD_REQUEST)
    }
}

func (ossServer *OssServer) DeleteHandler(w http.ResponseWriter, r *http.Request)  {
    ossRequest := NewOssRequest(r)
    ossRequest.ParseDelete()

    if ossRequest.Cmd == "" {
        panic("Cmd 空")
    }

    ossObject := OssObject{}
    ossBucket := OssBucket{}
    ossResponse := OssResponse{}
    switch ossRequest.Cmd {
    case DELETE_BUCKET:
        ossBucket.DeleteBucket(ossRequest.Bucket, w)
    case DELETE_BUCKET_LOGGING:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case DELETE_BUCKET_WEBSITE:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case DELETE_BUCKET_LIFECYCLE:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case DELETE_OBJECT:
        ossObject.DeleteObject(ossRequest.Bucket, ossRequest.Object, w)
    case DELETE_ABORT_MULTIPART_UPLOAD:
        ossResponse.ResponseError(w, NOT_IMPLEMENTED)
    case REQUEST_ERROR:
         ossResponse.ResponseError(w, BAD_REQUEST)
    default:
         ossResponse.ResponseError(w, BAD_REQUEST)
    }
}


func (ossServer *OssServer) OptionsHandler(w http.ResponseWriter, r *http.Request)  {
    ossResponse := OssResponse{}
    ossResponse.ResponseOptions(w)
}
