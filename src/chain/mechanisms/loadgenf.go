package mechanisms

import (
	"encoding/json"
	"io/ioutil"
)

func loadGenesisFromFile(filePath string) (block, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return block{}, err
	}

	var genesisData block
	err = json.Unmarshal(file, &genesisData)
	if err != nil {
		return block{}, err
	}

	return genesisData, nil
}