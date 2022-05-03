package encryption

import "errors"

type CryptoType uint32

// Errors.
const (
	errorUnknownCryptoType = `unknown crypto type`
	errorInvalidKeyLength  = `invalid key length`
)

// Crypto type.
const (
	AES128 CryptoType = 1 << (32 - 1 - iota)
	AES192
	AES256
)

var (
	keyLengthAESMap = map[CryptoType]AESKeyLength{
		AES128: keyLengthAES128,
		AES192: keyLengthAES192,
		AES256: keyLengthAES256,
	}
)

func MakeKey(t CryptoType) ([]byte, error) {
	if t <= AES256 {
		return makeAESKey(keyLengthAESMap[t])
	}
	return nil, errors.New(errorUnknownCryptoType)
}

func validateKey(t CryptoType, key []byte) error {
	var err = error(nil)
	switch t {
	case AES128:
		if len(key) != int(keyLengthAES128) {
			err = errors.New(errorInvalidKeyLength)
		}
	case AES192:
		if len(key) != int(keyLengthAES192) {
			err = errors.New(errorInvalidKeyLength)
		}
	case AES256:
		if len(key) != int(keyLengthAES256) {
			err = errors.New(errorInvalidKeyLength)
		}
	default:
		err = errors.New(errorUnknownCryptoType)
	}
	return err
}

func Encrypt(data []byte, t CryptoType, key []byte) ([]byte, error) {
	var ret []byte
	var err error

	// Validate key.
	err = validateKey(t, key)
	if err != nil {
		return nil, err
	}
	if t <= AES256 {
		ret, err = encryptAES(data, key)
	}
	return ret, err
}

func EncryptString(data string, t CryptoType, key []byte) ([]byte, error) {
	return Encrypt([]byte(data), t, key)
}

func Decrypt(data []byte, t CryptoType, key []byte) ([]byte, error) {
	var ret []byte
	var err error

	// Validate key.
	err = validateKey(t, key)
	if err != nil {
		return nil, err
	}
	if t <= AES256 {
		ret, err = decryptAES(data, key)
	}
	return ret, err
}

func DecryptString(data string, t CryptoType, key []byte) ([]byte, error) {
	return Decrypt([]byte(data), t, key)
}
