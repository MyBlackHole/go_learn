package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)
type AesTest struct {
	Key string
	Iv  string
}

// Encode 开始加密
func (a *AesTest) Encode(data string) (string, error) {
	_data := []byte(data)
	_key := []byte(a.Key)
	_iv := []byte(a.Iv)

	_data = a.PKCS7Padding(_data)
	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	return base64.StdEncoding.EncodeToString(_data), nil
}

// Decode 开始解密
func (a *AesTest) Decode(data string) (string, error) {
	_data, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	_key := []byte(a.Key)
	_iv := []byte(a.Iv)

	block, err := aes.NewCipher(_key)
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCDecrypter(block, _iv)
	mode.CryptBlocks(_data, _data)
	_data = a.PKCS7UnPadding(_data)

	return string(_data), nil
}
func (a *AesTest) PKCS7Padding(data []byte) []byte {
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
func (a *AesTest) PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
func main() {
	//data := "解密结果"
	data := "yDhlbcPUeHh2F/PpOUSN8AOCKm3Vnw3v5m60j+DwnhmZ8uBTdSvQKw0q5tP8XnEJxUFwXPp6tdKoxqS6WssWGUzMrELvRbYaiQj1xOQDNfjg8pPmxsFpbc8IIvR/Xn1FynA9ZqhHUGTpXYaFcP0Bg7XalxlwQMRLh9hJfa2vyB7AolslPGpDT5EJvttcgRwyPMWD7V7RjDPAe+n/Qwyxdw=="
	key := "3b9e61ed65ec555f"
	iv := "3b9e61ed65ec555f"
	aesTest := AesTest{
		Key: key,
		Iv:  iv,
	}
	if r, err := aesTest.Decode(data); err == nil {
		fmt.Printf("解密结果:%v\n", r)
	}
	//if r, err := aesTest.Encode(data); err == nil {
	//	fmt.Printf("加密结果:%v\n", r)
	//
	//	if r, err := aesTest.Decode(r); err == nil {
	//		fmt.Printf("解密结果:%v\n", r)
	//	}
	//}
}
