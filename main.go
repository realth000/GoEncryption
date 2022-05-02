package main

import (
	"GoEncryption/encryption"
	"fmt"
)

func main() {
	key, err := encryption.TestInitKey(encryption.AES256)

	c, err := encryption.Encrypt(key, "TestAbc123")
	if err != nil {
		fmt.Println(err)
		return
	}
	e, err := encryption.Decrypt(key, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", e)
}
