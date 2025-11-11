package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

// ForensicEvent represents a single file integrity event
type ForensicEvent struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
	FilePath  string `json:"file_path"`
	OldHash   string `json:"old_hash,omitempty"`
	NewHash   string `json:"new_hash,omitempty"`
}

// SaveForensicEvent appends forensic events to reports/forensic_log.json
func SaveForensicEvent(event ForensicEvent) {
	reportDir := "reports"
	reportFile := reportDir + "/forensic_log.json"
	os.MkdirAll(reportDir, 0755)

	var events []ForensicEvent
	if data, err := os.ReadFile(reportFile); err == nil && len(data) > 0 {
		_ = json.Unmarshal(data, &events)
	}

	events = append(events, event)
	data, _ := json.MarshalIndent(events, "", "  ")
	os.WriteFile(reportFile, data, 0644)

	// 🔐 Update integrity signature
	SaveSignature(reportFile)
}

// WatchDirectory monitors directory changes and updates logs/baseline
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
				SaveSignature(baselineFile)

				SaveForensicEvent(ForensicEvent{
					Timestamp: time.Now().Format(time.RFC3339),
					EventType: "ADDED",
					FilePath:  path,
					NewHash:   hash,
				})
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				newHash := GetFileHash(path)
				oldHash := baseline[path]
				if oldHash != newHash {
					log.Printf("[MODIFIED] %s\n  Old: %s\n  New: %s\n", path, oldHash, newHash)
					baseline[path] = newHash
					SaveBaseline(baseline, baselineFile)
					SaveSignature(baselineFile)

					SaveForensicEvent(ForensicEvent{
						Timestamp: time.Now().Format(time.RFC3339),
						EventType: "MODIFIED",
						FilePath:  path,
						OldHash:   oldHash,
						NewHash:   newHash,
					})
				}
			}

			if event.Op&fsnotify.Remove == fsnotify.Remove {
				log.Printf("[DELETED] %s\n", path)
				oldHash := baseline[path]
				delete(baseline, path)
				SaveBaseline(baseline, baselineFile)
				SaveSignature(baselineFile)

				SaveForensicEvent(ForensicEvent{
					Timestamp: time.Now().Format(time.RFC3339),
					EventType: "DELETED",
					FilePath:  path,
					OldHash:   oldHash,
				})
			}

		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}
