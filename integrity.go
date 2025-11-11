package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// ComputeFileHash returns the SHA256 hash of a given file
func ComputeFileHash(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// SaveSignature writes the hash to a .sig file
func SaveSignature(targetFile string) error {
	hash, err := ComputeFileHash(targetFile)
	if err != nil {
		return err
	}

	sigFile := targetFile + ".sig"
	return os.WriteFile(sigFile, []byte(hash), 0644)
}

// VerifySignature compares the file’s hash with the stored signature
func VerifySignature(targetFile string) (bool, error) {
	expectedHash, err := os.ReadFile(targetFile + ".sig")
	if err != nil {
		return false, fmt.Errorf("signature missing: %v", err)
	}

	currentHash, err := ComputeFileHash(targetFile)
	if err != nil {
		return false, err
	}

	return currentHash == string(expectedHash), nil
}
