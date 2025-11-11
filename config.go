package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config struct holds TraceLock configuration
type Config struct {
	MonitorDir string `json:"monitor_dir"`
	LogFile    string `json:"log_file"`
}

// LoadConfig loads settings from config.json, or applies defaults
func LoadConfig(path string) Config {
	var cfg Config

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("[!] config.json not found — using defaults.")
		cfg.MonitorDir = "./watched"
		cfg.LogFile = "./logs/tracelock.log"
		return cfg
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		fmt.Println("[!] Invalid config.json — using defaults.")
		cfg.MonitorDir = "./watched"
		cfg.LogFile = "./logs/tracelock.log"
	}

	return cfg
}
