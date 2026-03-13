package storage

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
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

// initDatabase initializes the database
func initDatabase() (*Database, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	if err := migrateLegacyData(homeDir); err != nil {
		return nil, err
	}

	// Create data directory
	dataDir := filepath.Join(homeDir, ".local", "share", "pock")
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	dbPath := filepath.Join(dataDir, "db.json")

	database := &Database{
		filePath: dbPath,
	}

	// Try to load existing database
	if err := database.load(); err != nil {
		// If file doesn't exist or is invalid, initialize with defaults
		database.Data = StorageData{
			CommandHistories: []CommandHistory{},
			SavedCommands:    []SavedCommand{},
			Settings:         DefaultSettings(),
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	scriptsDir := filepath.Join(homeDir, ".local", "share", "pock", "scripts")
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

// save writes the database to disk
func (db *Database) save() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := json.MarshalIndent(db.Data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(db.filePath, data, 0644)
}

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
