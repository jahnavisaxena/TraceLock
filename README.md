# 🛡️ FIMon — Real-Time File Integrity Monitoring Tool

**FIMon** is a lightweight, developer-friendly, real-time **File Integrity Monitoring (FIM)** tool written in Go.  
It continuously watches directories, computes **SHA-256 hashes**, and logs or alerts on any unauthorized file modifications.

---

##  Features

- 🔍 Real-time file monitoring using [`fsnotify`](https://github.com/fsnotify/fsnotify)
- 🧮 SHA-256 hash verification for integrity checks
- 🧾 Automatic baseline generation and updates
- 📜 Logging with timestamps and event type
- 🧠 Modular design for extension (email alerts, anomaly detection, DB storage)

---

##  Tech Stack

| Component | Tool / Package | Purpose |
|------------|----------------|----------|
| Language | Go (1.22+) | Core language |
| Monitoring | `fsnotify` | Real-time event watcher |
| Hashing | `crypto/sha256` | File integrity validation |
| Logging | Go `log` | Structured event logs |
| Storage | JSON baseline | Lightweight persistence |

---

##  Installation

```bash
# Clone repository
git clone https://github.com/jahnavisaxena/FIMon.git
cd FIMon

# Install dependencies
go mod tidy

# Build executable
go build -o fimon
```
##  Project Roadmap — FIMon Evolution Plan

FIMon is being developed as a modular, developer-friendly security agent that grows from a simple file integrity monitor into a lightweight host-based intrusion detection and response (HIDS/EDR) system.

| Phase | Name | Key Additions | Outcome |
|-------|------|---------------|----------|
| **1** | **Foundation (Core FIM)** | Real-time file watcher, SHA256 hash verification, baseline & logging | Detects unauthorized file modifications instantly |
| **2** | **Intelligence Layer** | Email / Telegram alerts, anomaly detection (burst monitoring) | Adds behavioral awareness and instant notifications |
| **3** |  **Persistence Layer** | SQLite audit logging, CLI queries, log rotation | Enables forensic history & reporting |
| **4** |  **Awareness Layer** | Process & user context (OSQuery / Falco), YARA scanning, Sigma-style rules | Understands “who” changed “what” — true HIDS |
| **5** |  **Visualization Layer** | Web dashboard (`net/http`), REST API, analytics graphs | Converts console tool into an interactive security dashboard |
| **6** |  **Integration Layer** | SIEM / Wazuh / Elastic export, Docker, systemd service | Enterprise-grade agent with centralized monitoring |

---

