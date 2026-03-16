//go:build !windows
// +build !windows

package storage

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

// acquireLock acquires an advisory file lock (Unix/Linux/macOS)
func (db *Database) acquireLock(lockPath string) error {
	// Create lock file
	lockFile, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	// Try to acquire exclusive lock with timeout
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		err = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
		if err == nil {
			db.lockFile = lockFile
			return nil
		}
		if err != syscall.EWOULDBLOCK {
			lockFile.Close()
			return err
		}
		// Wait before retry
		time.Sleep(100 * time.Millisecond)
	}

	lockFile.Close()
	return fmt.Errorf("failed to acquire lock after %d retries", maxRetries)
}

// releaseLock releases the advisory file lock
func (db *Database) releaseLock() error {
	if db.lockFile != nil {
		syscall.Flock(int(db.lockFile.Fd()), syscall.LOCK_UN)
		return db.lockFile.Close()
	}
	return nil
}
