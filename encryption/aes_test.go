package encryption

import (
	"bytes"
	"testing"
)

func TestAES256(t *testing.T) {
	key, err := MakeKey(AES256)
	if err != nil {
		t.Errorf("error making AES256 key:%s\n", err)
	}
	pOrig := "abcDEFINE"

	c, err := Encrypt([]byte(pOrig), AES256, key)
	if err != nil {
		t.Errorf("error encrypting AES-256:%s\n", err)
	}

	p, err := Decrypt(c, AES256, key)
	if err != nil {
		t.Errorf("error decrypting AES-256:%s\n", err)
	}

	if !bytes.Equal(p, []byte(pOrig)) {
		t.Errorf("error testing AES-256:plain test and cipher test not equal\n")
	}
}
