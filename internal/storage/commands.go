package storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CreateSavedCommandDB creates a new saved command
func CreateSavedCommandDB(cmd CreateSavedCommand) (*SavedCommand, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	newCommand := SavedCommand{
		ID:          uuid.New().String(),
		Name:        cmd.Name,
		Command:     cmd.Command,
		Description: cmd.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err = db.Update(func(data *StorageData) {
		data.SavedCommands = append(data.SavedCommands, newCommand)
	})

	if err != nil {
		return nil, err
	}

	return &newCommand, nil
}

// GetAllSavedCommands returns all saved commands
func GetAllSavedCommands() ([]SavedCommand, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var commands []SavedCommand
	db.Read(func(data *StorageData) {
		commands = make([]SavedCommand, len(data.SavedCommands))
		copy(commands, data.SavedCommands)
	})

	return commands, nil
}

// GetSavedCommandByName finds a command by name (case-insensitive)
func GetSavedCommandByName(name string) (*SavedCommand, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var found *SavedCommand
	db.Read(func(data *StorageData) {
		for _, cmd := range data.SavedCommands {
			if strings.EqualFold(cmd.Name, name) {
				found = &cmd
				return
			}
		}
	})

	return found, nil
}

// DeleteSavedCommand deletes a command by ID
func DeleteSavedCommand(id string) (bool, error) {
	db, err := GetDatabase()
	if err != nil {
		return false, err
	}

	var deleted bool
	err = db.Update(func(data *StorageData) {
		for i, cmd := range data.SavedCommands {
			if cmd.ID == id {
				data.SavedCommands = append(data.SavedCommands[:i], data.SavedCommands[i+1:]...)
				deleted = true
				return
			}
		}
	})

	return deleted, err
}

// UpdateSavedCommand updates a command by ID
func UpdateSavedCommand(id string, updates map[string]interface{}) (*SavedCommand, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var updated *SavedCommand
	err = db.Update(func(data *StorageData) {
		for i := range data.SavedCommands {
			if data.SavedCommands[i].ID == id {
				// Apply updates
				if name, ok := updates["name"].(string); ok {
					data.SavedCommands[i].Name = name
				}
				if command, ok := updates["command"].(string); ok {
					data.SavedCommands[i].Command = command
				}
				if description, ok := updates["description"].(string); ok {
					data.SavedCommands[i].Description = description
				}
				data.SavedCommands[i].UpdatedAt = time.Now()
				updated = &data.SavedCommands[i]
				return
			}
		}
	})

	if updated == nil {
		return nil, fmt.Errorf("command not found")
	}

	return updated, err
}
