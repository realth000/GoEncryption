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

type AESKeyLength int

const (
	keyLengthAES128 AESKeyLength = 16
	keyLengthAES192 AESKeyLength = 24
	keyLengthAES256 AESKeyLength = 32
)

func makeAESKey(length AESKeyLength) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func addAESPadding(data []byte) []byte {
	// TODO: Check when %= 0
	//paddingLength := (aes.BlockSize - len(data)%aes.BlockSize) % aes.BlockSize
	paddingLength := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, padding...)
}

func removeAESPadding(data []byte) []byte {
	paddingLength := data[len(data)-1]
	return data[:len(data)-int(paddingLength)]
}

func encryptAES(data []byte, key []byte) ([]byte, error) {

	data = addAESPadding(data)
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New(fmt.Sprintf("too short cipher text:len=%d", len(data)))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], data)
	return cipherText, nil
}

func decryptAES(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		fmt.Println(len(data))
		return nil, errors.New(fmt.Sprintf("too short cipher text:len=%d", len(data)))
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New(fmt.Sprintf("not regular cipher text length:len=%d", len(data)))
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return removeAESPadding(data), nil
}
