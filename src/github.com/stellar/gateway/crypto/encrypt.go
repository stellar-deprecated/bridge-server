package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/stellar/go-stellar-base/strkey"
)

func Encrypt(key string, message []byte) (cipherBase64, nonceBase64 string, err error) {
	keyBytes, err := strkey.Decode(strkey.VersionByteAccountID, key)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonceBytes := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonceBytes); err != nil {
		return
	}
	nonceBase64 = base64.StdEncoding.EncodeToString(nonceBytes)

	cipherBytes := aesgcm.Seal(nil, nonceBytes, message, nil)
	cipherBase64 = base64.StdEncoding.EncodeToString(cipherBytes)
	return
}

func Decrypt(key, cipherBase64, nonceBase64 string) (string, error) {
	keyBytes, err := strkey.Decode(strkey.VersionByteAccountID, key)
	if err != nil {
		return "", nil
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(cipherBase64)
	if err != nil {
		return "", err
	}

	nonceBytes, err := base64.StdEncoding.DecodeString(nonceBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonceBytes, cipherBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
