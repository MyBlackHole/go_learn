package emulator

func (fi FileInfo) ToObjectInfo(bucket, object string) ObjectInfo {
	objInfo := ObjectInfo{
		IsDir:   HasSuffix(object, SlashSeparator),
		Bucket:  bucket,
		Name:    object,
		Size:    fi.Size,
		ModTime: fi.ModTime,
		ETag:    "00000000000000000",
	}
	return objInfo
}
