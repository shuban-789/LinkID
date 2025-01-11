package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
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

func generateKeyPair() (string, string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(key), hex.EncodeToString(key), nil
}

func encrypt(data []byte, keyString string) ([]byte, error) {
	key, _ := hex.DecodeString(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

func decrypt(ciphertext []byte, keyString string) ([]byte, error) {
	key, _ := hex.DecodeString(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("Ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}