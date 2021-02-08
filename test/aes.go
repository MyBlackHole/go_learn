package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func main() {
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	iv, _ := hex.DecodeString("3b9e61ed65ec555f43f9fcb41d5dde3a")
	//plaintext := []byte("我是 password ")

	// -------- 加密开始---------
	file, err := os.Open("./config.json")
	data, _ := ioutil.ReadAll(file)
	plaintext, _ := base64.StdEncoding.DecodeString(string(data))

	plaintext = PKCS7Padding(plaintext)
	ciphertext := make([]byte, len(plaintext))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	fmt.Printf("%x\n", ciphertext)

	// ----------------解密开始---------

	mode = cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = PKCS7UnPadding(ciphertext)
	fmt.Printf("%s\n", ciphertext)
}
