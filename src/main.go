package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func help() {
	fmt.Println("Usage: ./linkid [OPTION1] [ARGUMENT1] ... [OPTIONn] [ARGUMENTn]\n")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -c, Create a new blockchain with the provided JSON file.")
	fmt.Println("  		-E, Save the output as JSON")
	fmt.Println("  -a, Access an existing blockchain with the provided ID and key.")
	fmt.Println("  		-E, Save the output as JSON")
	fmt.Println("  -A, Add a new block to an existing blockchain with the provided ID and key.")
	fmt.Println("")
	fmt.Println("Format:")
	fmt.Println("  ./linkid -c <GENESIS.json>")
	fmt.Println("  ./linkid -a <ID> <KEY>")
	fmt.Println("  ./linkid -A <BLOCK.json> <ID> <KEY>")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  ./linkid -c genesis.json")
	fmt.Println("  ./linkid -a 12345678 1234567890abcdef1234567890abcdef")
	fmt.Println("  ./linkid -A block.json 12345678 1234567890abcdef1234567890abcdef")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid command.")
		return
	}

	saveToFile := false
	if len(os.Args) > 2 && os.Args[2] == "-E" {
		saveToFile = true
	}

	switch os.Args[1] {
	case "-c":
		if len(os.Args) < 3 {
			fmt.Println("Error: Missing path to the input JSON file.")
			return
		}

		createCommand := os.Args[2]
		if saveToFile {
			createCommand = os.Args[3]
		}
		GenesisBlock, err := loadGenesisFromFile(createCommand)
		if err != nil {
			fmt.Println("Error loading genesis data:", err)
			return
		}

		GenesisBlock.Index = 0
		GenesisBlock.Time = time.Now().String()
		GenesisBlock.PreviousHash = ""
		GenesisBlock.CurrentHash = calculateHash(GenesisBlock)

		ChainID, err := generateChainID()
		if err != nil {
			fmt.Println("Error generating Chain ID:", err)
			return
		}

		TestChain := chain{
			ChainID:    ChainID,
			BlockCount: 1,
			Genesis:    GenesisBlock,
			Head:       GenesisBlock,
			Previous:   GenesisBlock,
			Chain:      []block{GenesisBlock},
		}

		publicKey, privateKey, err := generateKeyPair()
		if err != nil {
			fmt.Println("Error generating key pair:", err)
			return
		}

		keys := map[string]interface{}{
			"id":          ChainID,
			"public_key":  publicKey,
			"private_key": privateKey,
		}

		jsonData, err := json.MarshalIndent(keys, "", "  ")
		if err != nil {
			fmt.Println("Error converting keys to JSON:", err)
			return
		}
		
		if saveToFile {
			chainJSON, err := json.MarshalIndent(keys, "", "  ")
			if err != nil {
				fmt.Println("Error converting chain to JSON:", err)
				return
			}
			err = ioutil.WriteFile("output.json", chainJSON, 0644)
			if err != nil {
				fmt.Println("Error saving chain to output.json:", err)
			} else {
				fmt.Println("Chain saved to output.json.")
			}
			return
		}

		fmt.Println(string(jsonData))

		err = exportEncryptedChain(TestChain, privateKey)
		if err != nil {
			fmt.Println("Error exporting encrypted chain:", err)
			return
		}

	case "-a":
		if len(os.Args) < 4 {
			fmt.Println("Error: Missing chain ID or private key.")
			return
		}

		accessCommand := os.Args[2]
		key := os.Args[3]
		if saveToFile {
			accessCommand = os.Args[3]
			key = os.Args[4]
		}
		TargetChain, err := loadEncryptedChain(accessCommand, key)
		if err != nil {
			fmt.Println("Error loading encrypted chain:", err)
			return
		}

		if saveToFile {
			chainJSON, err := json.MarshalIndent(TargetChain, "", "  ")
			if err != nil {
				fmt.Println("Error converting chain to JSON:", err)
				return
			}
			err = ioutil.WriteFile("output.json", chainJSON, 0644)
			if err != nil {
				fmt.Println("Error saving chain to output.json:", err)
			} else {
				fmt.Println("Chain saved to output.json.")
			}
			return
		}

		jsonData, err := json.MarshalIndent(TargetChain, "", "  ")
		if err != nil {
			fmt.Println("Error converting chain to JSON:", err)
			return
		}

		fmt.Println(string(jsonData))

	case "-A":
		if len(os.Args) < 5 {
			fmt.Println("Error: Missing chain ID, block data file, or private key.")
			return
		}

		addBlockCommand := os.Args[2]
		blockChainID := os.Args[3]
		key := os.Args[4]

		TargetChain, err := loadEncryptedChain(blockChainID, key)
		if err != nil {
			fmt.Println("Error loading encrypted chain:", err)
			return
		}

		newBlockData, err := loadBlockFromFile(addBlockCommand)
		if err != nil {
			fmt.Println("Error loading block data:", err)
			return
		}

		addedBlock := generateBlock(TargetChain.Head, []interface{}{
			newBlockData.Index,
			newBlockData.Initials,
			newBlockData.Sex,
			newBlockData.Gender,
			newBlockData.Age,
			newBlockData.Height,
			newBlockData.Weight,
			newBlockData.BMI,
			newBlockData.Blood,
			newBlockData.Location,
			newBlockData.Prescriptions,
			newBlockData.Conditions,
			newBlockData.VisitLogs,
			newBlockData.History,
		})
		addBlockToChain(addedBlock, &TargetChain)

		if saveToFile {
			chainJSON, err := json.MarshalIndent(TargetChain, "", "  ")
			if err != nil {
				fmt.Println("Error converting chain to JSON:", err)
				return
			}
			err = ioutil.WriteFile("output.json", chainJSON, 0644)
			if err != nil {
				fmt.Println("Error saving chain to output.json:", err)
			} else {
				fmt.Println("Chain saved to output.json.")
			}
			return
		}

		err = exportEncryptedChain(TargetChain, key)
		if err != nil {
			fmt.Println("Error exporting encrypted chain:", err)
			return
		}

	default:
		help()
	}
}