package main

import (
	"encoding/xml"
	"os"
)

type ListBucketResult struct {
	XMLName      xml.Name `xml:"ListBucketResult"`
	Text         string   `xml:",chardata"`
	Xmlns        string   `xml:"xmlns,attr"`
	Name         string   `xml:"Name"`
	Prefix       string   `xml:"Prefix"`
	Marker       string   `xml:"Marker"`
	MaxKeys      string   `xml:"MaxKeys"`
	Delimiter    string   `xml:"Delimiter"`
	EncodingType string   `xml:"EncodingType"`
	NextMarker   string   `xml:"NextMarker"`
	IsTruncated  string   `xml:"IsTruncated"`
	Contents     []struct {
		Text         string `xml:",chardata"`
		Key          string `xml:"Key"`
		LastModified string `xml:"LastModified"`
		ETag         string `xml:"ETag"`
		Type         string `xml:"Type"`
		Size         string `xml:"Size"`
		StorageClass string `xml:"StorageClass"`
		Owner        struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"ID"`
			DisplayName string `xml:"DisplayName"`
		} `xml:"Owner"`
	} `xml:"Contents"`
}

var XMLNS = "http://doc.oss-cn-hangzhou.aliyuncs.com"

func main() {
	data := &ListBucketResult{}
    data.Xmlns = XMLNS
	// data.XMLName = xml.Name{
	// 	Space: XMLNS,
 //        Local: XMLNS,
	// }
	data.Name = "wdg1"
	data.Marker = "/"
	data.Prefix = "wdg1"
	data.MaxKeys = "1000"
	data.Delimiter = ""
	data.EncodingType = "url"
    out, err := xml.Marshal(data)
    if err != nil {
        panic(err)
    }

    err = os.WriteFile("./list_bucket_result.xml", out, os.ModePerm)
    if err != nil {
        panic(err)
    }
}
