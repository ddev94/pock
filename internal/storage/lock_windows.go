//go:build windows
// +build windows

package storage

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procLockFileEx   = modkernel32.NewProc("LockFileEx")
	procUnlockFileEx = modkernel32.NewProc("UnlockFileEx")
)

const (
	lockfileExclusiveLock   = 0x00000002
	lockfileFailImmediately = 0x00000001
)

// acquireLock acquires an advisory file lock (Windows)
func (db *Database) acquireLock(lockPath string) error {
	// Create lock file
	lockFile, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	// Try to acquire exclusive lock with timeout
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		err = lockFileEx(syscall.Handle(lockFile.Fd()), lockfileExclusiveLock|lockfileFailImmediately, 1, 0, &syscall.Overlapped{})
		if err == nil {
			db.lockFile = lockFile
			return nil
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
		unlockFileEx(syscall.Handle(db.lockFile.Fd()), 1, 0, &syscall.Overlapped{})
		return db.lockFile.Close()
	}
	return nil
}

// lockFileEx wraps the Windows LockFileEx function
func lockFileEx(handle syscall.Handle, flags, lockLow, lockHigh uint32, overlapped *syscall.Overlapped) error {
	r1, _, err := procLockFileEx.Call(
		uintptr(handle),
		uintptr(flags),
		uintptr(0),
		uintptr(lockLow),
		uintptr(lockHigh),
		uintptr(unsafe.Pointer(overlapped)),
	)
	if r1 == 0 {
		return err
	}
	return nil
}

// unlockFileEx wraps the Windows UnlockFileEx function
func unlockFileEx(handle syscall.Handle, lockLow, lockHigh uint32, overlapped *syscall.Overlapped) error {
	r1, _, err := procUnlockFileEx.Call(
		uintptr(handle),
		uintptr(0),
		uintptr(lockLow),
		uintptr(lockHigh),
		uintptr(unsafe.Pointer(overlapped)),
	)
	if r1 == 0 {
		return err
	}
	return nil
}
