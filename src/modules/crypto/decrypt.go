package modules

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

func decrypt(ciphertext []byte, keyString string) ([]byte, error) {
	key, _ := hex.DecodeString(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("Ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}