package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CreateBaseline(dir string, baselinePath string) map[string]string {
	baseline := make(map[string]string)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		baseline[path] = GetFileHash(path)
		return nil
	})

	data, _ := json.MarshalIndent(baseline, "", "  ")
	os.WriteFile(baselinePath, data, 0644)
	fmt.Println("[+] Baseline created at:", baselinePath)
	return baseline
}

func LoadBaseline(baselinePath string) map[string]string {
	baseline := make(map[string]string)
	file, err := os.Open(baselinePath)
	if err != nil {
		return baseline
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&baseline)
	return baseline
}
