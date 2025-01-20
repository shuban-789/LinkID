package main

import (
    "encoding/pem"
    "crypto/rsa"
    "crypto/rand"
    "crypto/sha256"
    "crypto/x509"
    "encoding/hex"
    "errors"
    "fmt"
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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Error generating private key: %v\n", err)
		return "", "", err
	}

	publicKey := &privateKey.PublicKey

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Printf("Error marshaling public key: %v\n", err)
		return "", "", err
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return string(privateKeyPEM), string(publicKeyPEM), nil
}

func encrypt(data []byte, publicKeyString string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyString))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key format")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("Not an RSA public key")
	}

	hash := sha256.New()

	encryptedData, err := rsa.EncryptOAEP(hash, rand.Reader, rsaPub, data, nil)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

func decrypt(ciphertext []byte, privateKeyString string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKeyString))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid private key format")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hash := sha256.New()

	decryptedData, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}