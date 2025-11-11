package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("🛡️  FIMon — Digital Forensic File Integrity Tool")
	fmt.Println("------------------------------------------------")

	cfg := LoadConfig("config.json")

	os.MkdirAll("logs", 0755)
	logFile, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Cannot open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	baselineFile := "baseline.json"

	if _, err := os.Stat(baselineFile); os.IsNotExist(err) {
		CreateBaseline(cfg.MonitorDir, baselineFile)
	}
	baseline := LoadBaseline(baselineFile)

	go WatchDirectory(cfg, baseline, baselineFile)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	<-done
	fmt.Println("\n🛑 Monitoring stopped.")
}
