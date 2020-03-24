// aes
package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AES struct {
	Key     string
	Origin  string
	Crypted string
}

// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
var key string = "&^3092$ks*&2dpx("

func (this *AES) init() {
	if len(this.Key) == 0 {
		this.Key = key
	}
}

func (this *AES) Encrypt() error {
	this.init()
	tmpKey := []byte(this.Key)
	block, err := aes.NewCipher(tmpKey)
	if err != nil {
		return err
	}

	blockSize := block.BlockSize()
	tmpOrigin := this.pkcs5Padding([]byte(this.Origin), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, tmpKey[:blockSize])
	tmpCrypted := make([]byte, len(tmpOrigin))
	blockMode.CryptBlocks(tmpCrypted, tmpOrigin)
	this.Crypted = string(tmpCrypted)
	return nil
}

func (this *AES) CryptedToString() string {
	return base64.RawURLEncoding.EncodeToString([]byte(this.Crypted))
}

func (this *AES) Decrypt() error {
	this.init()
	tmpKey := []byte(this.Key)
	block, err := aes.NewCipher(tmpKey)
	if err != nil {
		return err
	}

	tmpCrypted, err := base64.RawURLEncoding.DecodeString(this.Crypted) //[]byte(this.Crypted)
	if err != nil {
		return err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, tmpKey[:blockSize])
	tmpOrigin := make([]byte, len(tmpCrypted))
	blockMode.CryptBlocks(tmpOrigin, tmpCrypted)
	this.Origin = string(this.pkcs5UnPadding(tmpOrigin))
	return nil
}

func (this *AES) pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func (this *AES) pkcs5UnPadding(originData []byte) []byte {
	length := len(originData)
	unpadding := int(originData[length-1])
	return originData[:(length - unpadding)]
}
