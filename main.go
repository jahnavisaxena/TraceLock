package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory_path>")
		return
	}

	path := os.Args[1]
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatalf("Error creating watcher: %v", err)
	}
	defer watcher.Close()

	err = watcher.Add(path)
	if err != nil {
		logger.Fatalf("Error watching directory: %v", err)
	}

	logger.Infof("Monitoring directory: %s", path)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			logger.Infof("File event: %s %s", event.Op, event.Name)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Errorf("Error: %v", err)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

