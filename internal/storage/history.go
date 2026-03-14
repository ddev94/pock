package storage

import (
	"time"

	"github.com/google/uuid"
)

// CreateCommandHistory creates a new command history entry
func CreateCommandHistory(commandName, commandText, status, log string, executionTime int64) (*CommandHistory, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	history := CommandHistory{
		ID:            uuid.New().String(),
		CommandName:   commandName,
		CommandText:   commandText,
		Date:          time.Now(),
		Status:        status,
		Log:           log,
		ExecutionTime: executionTime,
	}

	err = db.Update(func(data *StorageData) {
		data.CommandHistories = append(data.CommandHistories, history)
	})

	if err != nil {
		return nil, err
	}

	return &history, nil
}

// GetCommandHistory returns all command history entries
func GetCommandHistory(limit int) ([]CommandHistory, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var histories []CommandHistory
	db.Read(func(data *StorageData) {
		// Get all histories
		allHistories := data.CommandHistories

		// If limit is specified and less than total, get the last N entries
		if limit > 0 && limit < len(allHistories) {
			start := len(allHistories) - limit
			histories = make([]CommandHistory, limit)
			copy(histories, allHistories[start:])
		} else {
			histories = make([]CommandHistory, len(allHistories))
			copy(histories, allHistories)
		}
	})

	return histories, nil
}

// GetCommandHistoryByName returns history entries for a specific command
func GetCommandHistoryByName(commandName string, limit int) ([]CommandHistory, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	var histories []CommandHistory
	db.Read(func(data *StorageData) {
		// Filter histories by command name
		var filtered []CommandHistory
		for _, h := range data.CommandHistories {
			if h.CommandName == commandName {
				filtered = append(filtered, h)
			}
		}

		// If limit is specified and less than total, get the last N entries
		if limit > 0 && limit < len(filtered) {
			start := len(filtered) - limit
			histories = make([]CommandHistory, limit)
			copy(histories, filtered[start:])
		} else {
			histories = make([]CommandHistory, len(filtered))
			copy(histories, filtered)
		}
	})

	return histories, nil
}

// GetCommandStats returns statistics for a command.
func GetCommandStats(commandName string) (*CommandStats, error) {
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}

	stats := &CommandStats{
		TotalRuns:        0,
		SuccessfulRuns:   0,
		FailedRuns:       0,
		LastRun:          "",
		AvgExecutionTime: 0,
	}

	db.Read(func(data *StorageData) {
		var totalExecTime int64
		var count int64

		for _, h := range data.CommandHistories {
			if h.CommandName == commandName {
				stats.TotalRuns++
				if h.Status == "success" {
					stats.SuccessfulRuns++
				} else {
					stats.FailedRuns++
				}
				stats.LastRun = h.Date.Format(time.RFC3339)
				if h.ExecutionTime > 0 {
					totalExecTime += h.ExecutionTime
					count++
				}
			}
		}

		if count > 0 {
			stats.AvgExecutionTime = totalExecTime / count
		}
	})

	return stats, nil
}

// ClearCommandHistory clears all command history
func ClearCommandHistory() error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	return db.Update(func(data *StorageData) {
		data.CommandHistories = []CommandHistory{}
	})
}

// ClearCommandHistoryByName clears history for a specific command
func ClearCommandHistoryByName(commandName string) error {
	db, err := GetDatabase()
	if err != nil {
		return err
	}

	return db.Update(func(data *StorageData) {
		var filtered []CommandHistory
		for _, h := range data.CommandHistories {
			if h.CommandName != commandName {
				filtered = append(filtered, h)
			}
		}
		data.CommandHistories = filtered
	})
}
