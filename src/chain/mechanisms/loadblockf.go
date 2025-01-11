package mechanisms

import (
	"encoding/json"
	"io/ioutil"
)

func loadBlockFromFile(filePath string) (block, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return block{}, err
	}

	var newBlock block
	err = json.Unmarshal(file, &newBlock)
	if err != nil {
		return block{}, err
	}

	return newBlock, nil
}