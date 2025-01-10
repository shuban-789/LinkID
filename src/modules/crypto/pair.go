package modules

import (
	"crypto/rand"
	"encoding/hex"
)

func generateKeyPair() (string, string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(key), hex.EncodeToString(key), nil
}