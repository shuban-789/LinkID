package main

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func generateBlock(previousBlock block, data []interface{}) block {
	NewBlock := block{
		Index:        	data[0].(int),
		Initials:     	data[1].(string),
		Sex:          	data[2].(string),
		Gender:       	data[3].(string),
		Age:          	data[4].(int),
		Height:       	data[5].(float32),
		Weight:       	data[6].(float32),
		BMI:          	data[7].(float32),
		Blood:        	data[8].(string),
		Time:         	time.Now().String(),
		Location:     	data[9].(string),
		PreviousHash: 	previousBlock.CurrentHash,
		Prescriptions:  data[10].([]string),
		Conditions:   	data[11].([]string),
		VisitLogs:    	data[12].([]string),
		History:      	data[13].([]string),
	}

	NewBlock.CurrentHash = calculateHash(NewBlock)
	return NewBlock
}

func addBlockToChain(AddedBlock block, TargetChain *chain) {
	TargetChain.Chain = append(TargetChain.Chain, AddedBlock)
	TargetChain.Previous = TargetChain.Head
	TargetChain.Head = AddedBlock
	TargetChain.BlockCount++
}

func mineBlock(previousBlock block, data []interface{}, difficulty int) block {
	var nonce int
	var NewBlock block

	for {
		NewBlock = block{
			Index:        	data[0].(int),
			Initials:     	data[1].(string),
			Sex:          	data[2].(string),
			Gender:       	data[3].(string),
			Age:          	data[4].(int),
			Height:       	data[5].(float32),
			Weight:       	data[6].(float32),
			BMI:          	data[7].(float32),
			Blood:        	data[8].(string),
			Time:         	time.Now().String(),
			Location:     	data[9].(string),
			PreviousHash: 	previousBlock.CurrentHash,
			Prescriptions:  data[10].([]string),
			Conditions:   	data[11].([]string),
			VisitLogs:    	data[12].([]string),
			History:      	data[13].([]string),
		}

		NewBlock.CurrentHash = calculateHash(NewBlock)

		if NewBlock.CurrentHash[:difficulty] == string(make([]byte, difficulty)) {
			break
		}
		nonce++
	}

	return NewBlock
}

func getBlockByHash(TargetChain chain, hash string) (block, bool) {
	for _, b := range TargetChain.Chain {
		if b.CurrentHash == hash {
			return b, true
		}
	}
	return block{}, false
}

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