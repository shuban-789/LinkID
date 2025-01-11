package init

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
)

func generateChainID() (int, error) {
	dir := "./records"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return 0, err
	}

	max := big.NewInt(90000000)
	min := 10000000

	for {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return 0, err
		}
		chainID := int(n.Int64()) + min
		encFilePath := filepath.Join(dir, strconv.Itoa(chainID) + ".enc")
		if _, err := os.Stat(encFilePath); os.IsNotExist(err) {
			return chainID, nil
		}
	}
}