
package main 
import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	MonitorDir string `json:"monitor_dir"`
	LogFile    string `json:"log_file"`
}

func LoadConfig() (Config, error) {
	var cfg Config

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("  No config.json found — using defaults.")
		cfg.MonitorDir = "./watched"
		cfg.LogFile = "./logs/fim.log"
		return cfg, nil
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

