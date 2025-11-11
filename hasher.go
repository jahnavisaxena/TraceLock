package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// GetFileHash generates a SHA-256 hash for the given file path
func GetFileHash(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("[ERROR] Unable to open file for hashing: %s | %v", path, err)
		return ""
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Printf("[ERROR] Unable to compute hash for: %s | %v", path, err)
		return ""
	}

	return hex.EncodeToString(hash.Sum(nil))
}
