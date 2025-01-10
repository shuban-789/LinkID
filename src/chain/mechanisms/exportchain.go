package mechanisms

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func exportEncryptedChain(TargetChain chain, key string) error {
	dir := "./records"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	filePath := filepath.Join(dir, strconv.Itoa(TargetChain.ChainID) + ".enc")

	chainJSON, err := json.MarshalIndent(TargetChain, "", "  ")
	if err != nil {
		return err
	}

	encryptedData, err := encrypt(chainJSON, key)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, encryptedData, 0644)
	if err != nil {
		return err
	}

	return nil
}