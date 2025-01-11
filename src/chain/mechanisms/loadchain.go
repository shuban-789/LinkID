package mechanisms

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
)

func loadEncryptedChain(id string, key string) (chain, error) {
	filePath := filepath.Join("./records", id+".enc")

	encryptedData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return chain{}, err
	}

	decryptedData, err := decrypt(encryptedData, key)
	if err != nil {
		return chain{}, err
	}

	var TargetChain chain
	err = json.Unmarshal(decryptedData, &TargetChain)
	if err != nil {
		return chain{}, err
	}

	return TargetChain, nil
}