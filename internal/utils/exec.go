package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// ExecResult represents the result of command execution
type ExecResult struct {
	Success       bool
	Output        string
	Error         string
	ExitCode      int
	ExecutionTime int64 // in milliseconds
}

// ExecuteCommand executes a shell command and returns the result
func ExecuteCommand(command string) *ExecResult {
	startTime := time.Now()

	// Get the user's shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	// Create the command
	cmd := exec.Command(shell, "-c", command)

	// Set up output buffers
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	executionTime := time.Since(startTime).Milliseconds()

	result := &ExecResult{
		Output:        stdout.String(),
		Error:         stderr.String(),
		ExecutionTime: executionTime,
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.ExitCode = 1
		}
		result.Success = false
	} else {
		result.ExitCode = 0
		result.Success = true
	}

	return result
}

// ExecuteCommandInteractive executes a command with interactive I/O
// It captures output while still displaying it to the user
func ExecuteCommandInteractive(command string) *ExecResult {
	startTime := time.Now()

	// Get the user's shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	// Create the command
	cmd := exec.Command(shell, "-c", command)

	// Set up output buffers that also write to stdout/stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	// Run the command
	err := cmd.Run()

	executionTime := time.Since(startTime).Milliseconds()

	result := &ExecResult{
		Output:        stdout.String(),
		Error:         stderr.String(),
		ExecutionTime: executionTime,
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.ExitCode = 1
		}
		result.Success = false
	} else {
		result.ExitCode = 0
		result.Success = true
	}

	return result
}

// IsScriptFile checks if the command is a script file
func IsScriptFile(command string) bool {
	// Check if it looks like a file path
	if strings.Contains(command, "/") ||
		strings.HasSuffix(command, ".sh") ||
		strings.HasSuffix(command, ".bash") {
		// Check if file exists
		if _, err := os.Stat(command); err == nil {
			return true
		}
	}
	return false
}

// FormatDuration formats a duration in milliseconds to a human-readable string
func FormatDuration(ms int64) string {
	if ms < 1000 {
		return fmt.Sprintf("%dms", ms)
	} else if ms < 60000 {
		return fmt.Sprintf("%.2fs", float64(ms)/1000)
	} else {
		minutes := ms / 60000
		seconds := (ms % 60000) / 1000
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
}
