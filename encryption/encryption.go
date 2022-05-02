package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

type AESKeyLength uint32

const (
	AES128 AESKeyLength = 16
	AES192 AESKeyLength = 24
	AES256 AESKeyLength = 32
)

func TestInitKey(length AESKeyLength) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func addPadding(data []byte) []byte {
	// TODO: Check when %= 0
	//paddingLength := (aes.BlockSize - len(data)%aes.BlockSize) % aes.BlockSize
	paddingLength := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, padding...)
}

func removePadding(data []byte) []byte {
	paddingLength := data[len(data)-1]
	return data[:len(data)-int(paddingLength)]
}

func Encrypt(key []byte, data string) ([]byte, error) {
	plainText := []byte(data)

	plainText = addPadding(plainText)
	if len(plainText)%aes.BlockSize != 0 {
		return nil, errors.New("not regular plain text length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	return cipherText, nil
}

func Decrypt(key []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		fmt.Println(len(data))
		return nil, errors.New("too short cipher text")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("not regular cipher text length")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return removePadding(data), nil
}
