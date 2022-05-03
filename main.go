package main

import (
	"C"
	"GoEncryption/encryption"
	"encoding/hex"
	"fmt"
)

const (
	testKey = `148cd21bb34af8db7d410201de6ec019058cfa62daecf3b69c481b6248765252`
)

//export MakeAES256Key
func MakeAES256Key() *C.char {
	key, err := encryption.MakeKey(encryption.AES256)
	// TODO: Handle error.
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return C.CString(string(key))
}

//export GoEncrypt
func GoEncrypt(data *C.char, keyByte *C.char) *C.char {
	c, err := encryption.Encrypt([]byte(C.GoString(data)), encryption.AES256, []byte(C.GoString(keyByte)))
	// TODO: Handle error.
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return C.CString(string(c))
}

//export GoDecrypt
func GoDecrypt(data *C.char, keyByte *C.char) *C.char {
	p, err := encryption.Decrypt([]byte(C.GoString(data)), encryption.AES256, []byte(C.GoString(keyByte)))
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
