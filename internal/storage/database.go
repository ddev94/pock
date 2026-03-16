package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	db   *Database
	once sync.Once
)

// Database represents the JSON file database
type Database struct {
	mu       sync.RWMutex
	filePath string
	lockFile *os.File
	Data     StorageData
}

// GetDatabase returns the singleton database instance
func GetDatabase() (*Database, error) {
	var err error
	once.Do(func() {
		db, err = initDatabase()
	})
	return db, err
}

// getDataDir returns the data directory path based on OS and environment variable
func getDataDir() (string, error) {
	// Check for override environment variable
	if customDir := os.Getenv("POCK_DATA_DIR"); customDir != "" {
		return customDir, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Platform-specific paths
	switch runtime.GOOS {
	case "darwin":
		// macOS: ~/Library/Application Support/pock
		return filepath.Join(homeDir, "Library", "Application Support", "pock"), nil
	case "windows":
		// Windows: %AppData%\pock
		appData := os.Getenv("AppData")
		if appData != "" {
			return filepath.Join(appData, "pock"), nil
		}
		return filepath.Join(homeDir, "AppData", "Roaming", "pock"), nil
	default:
		// Linux/Unix: XDG Base Directory specification
		xdgDataHome := os.Getenv("XDG_DATA_HOME")
		if xdgDataHome != "" {
			return filepath.Join(xdgDataHome, "pock"), nil
		}
		return filepath.Join(homeDir, ".local", "share", "pock"), nil
	}
}

// initDatabase initializes the database
func initDatabase() (*Database, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	if err := migrateLegacyData(homeDir); err != nil {
		return nil, err
	}

	// Get data directory based on OS and environment
	dataDir, err := getDataDir()
	if err != nil {
		return nil, err
	}

	// Create data directory
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(dataDir, "db.json")
	lockPath := filepath.Join(dataDir, "db.lock")

	database := &Database{
		filePath: dbPath,
	}

	// Acquire file lock
	if err := database.acquireLock(lockPath); err != nil {
		return nil, fmt.Errorf("failed to acquire lock: %w", err)
	}

	// Try to load existing database
	if err := database.load(); err != nil {
		// If file doesn't exist or is invalid, initialize with defaults
		database.Data = StorageData{
			CommandHistories: []CommandHistory{},
			SavedCommands:    []SavedCommand{},
		}
		// Save the initial data
		if err := database.save(); err != nil {
			return nil, err
		}
	}

	return database, nil
}

func migrateLegacyData(homeDir string) error {
	legacyPath := filepath.Join(homeDir, ".local", "share", "hish", "db.json")
	newPath := filepath.Join(homeDir, ".local", "share", "pock", "db.json")

	if _, err := os.Stat(newPath); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	legacyFile, err := os.Open(legacyPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer legacyFile.Close()

	if err := os.MkdirAll(filepath.Dir(newPath), 0755); err != nil {
		return err
	}

	newFile, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, legacyFile)
	return err
}

// GetScriptsDir returns (and creates if needed) the directory where pock stores managed script files.
func GetScriptsDir() (string, error) {
	dataDir, err := getDataDir()
	if err != nil {
		return "", err
	}
	scriptsDir := filepath.Join(dataDir, "scripts")
	if err := os.MkdirAll(scriptsDir, 0755); err != nil {
		return "", err
	}
	return scriptsDir, nil
}

// load reads the database from disk
func (db *Database) load() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := os.ReadFile(db.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &db.Data)
}

// save writes the database to disk using atomic write
func (db *Database) save() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := json.MarshalIndent(db.Data, "", "  ")
	if err != nil {
		return err
	}

	// Atomic write: write to temp file, sync, then rename
	tmpPath := db.filePath + ".tmp"

	// Write to temporary file
	tmpFile, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	// Write data
	if _, err := tmpFile.Write(data); err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		return err
	}

	// Sync to ensure data is written to disk
	if err := tmpFile.Sync(); err != nil {
		tmpFile.Close()
		os.Remove(tmpPath)
		return err
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpPath)
		return err
	}

	// Atomic rename
	return os.Rename(tmpPath, db.filePath)
}

// acquireLock and releaseLock are implemented in platform-specific files:
// - lock_unix.go for Unix/Linux/macOS
// - lock_windows.go for Windows

// Update updates the database with a function
func (db *Database) Update(fn func(*StorageData)) error {
	db.mu.Lock()
	fn(&db.Data)
	db.mu.Unlock()
	return db.save()
}

// Read performs a read operation with the database
func (db *Database) Read(fn func(*StorageData)) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	fn(&db.Data)
}
