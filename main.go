package main

import (
	"GoEncryption/encryption"
	"encoding/hex"
	"fmt"
)

const (
	testKey = `148cd21bb34af8db7d410201de6ec019058cfa62daecf3b69c481b6248765252`
)

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
