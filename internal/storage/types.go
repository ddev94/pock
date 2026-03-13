package storage

import "time"

// CommandHistory represents a command execution history entry
type CommandHistory struct {
	ID            string    `json:"id"`
	CommandName   string    `json:"commandName"`
	CommandText   string    `json:"commandText"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"` // "success" or "failure"
	Log           string    `json:"log"`
	ExecutionTime int64     `json:"executionTime,omitempty"` // in milliseconds
}

// SavedCommand represents a saved command
type SavedCommand struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// CreateSavedCommand represents the data needed to create a new command
type CreateSavedCommand struct {
	Name        string `json:"name"`
	Command     string `json:"command"`
	Description string `json:"description,omitempty"`
}

// Settings represents user settings
type Settings struct {
	ListLayout string `json:"listLayout"` // "table" or "simple"
	DateFormat string `json:"dateFormat"` // "relative", "locale", or "iso"
}

// DefaultSettings returns the default settings
func DefaultSettings() Settings {
	return Settings{
		ListLayout: "table",
		DateFormat: "locale",
	}
}

// MarketplaceCommand represents a command from the marketplace
type MarketplaceCommand struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Description string   `json:"description,omitempty"`
	Author      string   `json:"author,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Version     string   `json:"version,omitempty"`
	Downloads   int      `json:"downloads,omitempty"`
	PublishedAt string   `json:"publishedAt,omitempty"`
}

// ExportedCommand represents a command for export
type ExportedCommand struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Description string   `json:"description,omitempty"`
	Author      string   `json:"author,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Version     string   `json:"version,omitempty"`
}

// StorageData represents the entire database structure
type StorageData struct {
	CommandHistories []CommandHistory `json:"commandHistories"`
	SavedCommands    []SavedCommand   `json:"savedCommands"`
	Settings         Settings         `json:"settings"`
}
