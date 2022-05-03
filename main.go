package main

import (
	"C"
	"GoEncryption/encryption"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	testKey = `148cd21bb34af8db7d410201de6ec019058cfa62daecf3b69c481b6248765252`
)

func base64Encode(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	return dst
}

func base64Decode(data []byte) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dst, data)
	if err != nil {
		fmt.Printf("base64 decode error:%s(len=%d)\n", err, n)
		return nil
	}
	return dst[:len(dst)-1]
}

//export MakeAES256KeyToBase64
func MakeAES256KeyToBase64() *C.char {
	key, err := encryption.MakeKey(encryption.AES256)
	// TODO: Handle error.
	if err != nil {
		fmt.Println(err)
		return nil
	}
	key = base64Encode(key)
	return C.CString(string(key))
}

//export GoEncryptToBase64
func GoEncryptToBase64(data *C.char, key *C.char) *C.char {
	keyRaw := base64Decode([]byte(C.GoString(key)))
	c, err := encryption.Encrypt([]byte(C.GoString(data)), encryption.AES256, keyRaw)
	// TODO: Handle error.
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return C.CString(string(base64Encode(c)))
}

//export GoDecryptToBase64
func GoDecryptToBase64(data *C.char, key *C.char) *C.char {
	keyRaw := base64Decode([]byte(C.GoString(key)))
	dataRaw := base64Decode([]byte(C.GoString(data)))
	p, err := encryption.Decrypt(dataRaw, encryption.AES256, keyRaw)
	// TODO: Handle error.
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return C.CString(string(p))
}

func main() {
	keyByte, _ := hex.DecodeString(testKey)
	c, err := encryption.EncryptString("TestAbc123", encryption.AES256, keyByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	p, err := encryption.Decrypt(c, encryption.AES256, keyByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\"%s\"\n", p)
}
