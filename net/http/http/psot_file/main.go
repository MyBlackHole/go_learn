package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "8.134.106.141:8068/ser/sva/get_speech_voice_performance"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("key", "f7ed5fcd-222f-5dd9-bf90-166d4b848264")
	_ = writer.WriteField("secret", "d8d1f74808dc8310f56edb8366b70d0fc036531a6622fecfff0ece09ab501277")
	_ = writer.WriteField("speech_text", "就是下雨也去。")
	_ = writer.WriteField("no_punc_word_list_str", "[\"就是\", \"下雨\", \"也\", \"去\"]")
	file, errFile5 := os.Open("/home/black/Documents/Code/ZFDM/go/go_test/audios/audio00.wav")
	defer file.Close()
	part5,
		errFile5 := writer.CreateFormFile("file", filepath.Base("/home/black/Documents/Code/ZFDM/go/go_test/audios/audio00.wav"))
	_, errFile5 = io.Copy(part5, file)
	if errFile5 != nil {
		fmt.Println(errFile5)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
