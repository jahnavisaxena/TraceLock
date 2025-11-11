package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

func WatchDirectory(cfg Config, baseline map[string]string, baselineFile string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(cfg.MonitorDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("[🔍] Watching directory: %s\n", cfg.MonitorDir)

	for {
		select {
		case event := <-watcher.Events:
			path := event.Name

			if event.Op&fsnotify.Create == fsnotify.Create {
				hash := GetFileHash(path)
				baseline[path] = hash
				log.Printf("[ADDED] %s | Hash: %s\n", path, hash)
				SaveBaseline(baseline, baselineFile)
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				newHash := GetFileHash(path)
				oldHash := baseline[path]
				if oldHash != newHash {
					log.Printf("[MODIFIED] %s\n  Old: %s\n  New: %s\n", path, oldHash, newHash)
					baseline[path] = newHash
					SaveBaseline(baseline, baselineFile)
				}
			}

			if event.Op&fsnotify.Remove == fsnotify.Remove {
				log.Printf("[DELETED] %s\n", path)
				delete(baseline, path)
				SaveBaseline(baseline, baselineFile)
			}

		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}

// SaveBaseline updates the baseline.json file after changes
func SaveBaseline(baseline map[string]string, baselineFile string) {
	data, _ := json.MarshalIndent(baseline, "", "  ")
	os.WriteFile(baselineFile, data, 0644)
}
