package mechanisms

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func calculateHash(b block) string {
	var BlockData = []string{
		strconv.Itoa(b.Index),
		b.Initials,
		b.Sex,
		b.Gender,
		strconv.Itoa(b.Age),
		b.Time,
		strconv.FormatFloat(float64(b.Height), 'f', -1, 32),
		strconv.FormatFloat(float64(b.Weight), 'f', -1, 32),
		strconv.FormatFloat(float64(b.BMI), 'f', -1, 32),
		b.Blood,
		b.Location,
		b.PreviousHash,
	}

	var record string
	for _, data := range BlockData {
		record += data
	}

	for _, med := range b.Prescriptions {
		record += med
	}

	for _, cond := range b.Conditions {
		record += cond
	}

	for _, visit := range b.VisitLogs {
		record += visit
	}

	for _, hist := range b.History {
		record += hist
	}

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}