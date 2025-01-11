package mechanisms

import (
	"time"
)

func mineBlock(previousBlock block, data []interface{}, difficulty int) block {
	var nonce int
	var NewBlock block

	for {
		NewBlock = block{
			Index:        data[0].(int),
			Initials:     data[1].(string),
			Sex:          data[2].(string),
			Gender:       data[3].(string),
			Age:          data[4].(int),
			Height:       data[5].(float32),
			Weight:       data[6].(float32),
			BMI:          data[7].(float32),
			Blood:        data[8].(string),
			Time:         time.Now().String(),
			Location:     data[9].(string),
			PreviousHash: previousBlock.CurrentHash,
			Prescriptions:  data[10].([]string),
			Conditions:   data[11].([]string),
			VisitLogs:    data[12].([]string),
			History:      data[13].([]string),
		}

		NewBlock.CurrentHash = calculateHash(NewBlock)

		if NewBlock.CurrentHash[:difficulty] == string(make([]byte, difficulty)) {
			break
		}
		nonce++
	}

	return NewBlock
}