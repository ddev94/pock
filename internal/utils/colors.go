package utils

import "github.com/fatih/color"

var (
	// Success colors
	Green  = color.New(color.FgGreen).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Blue   = color.New(color.FgBlue).SprintFunc()
	
	// Warning/Error colors
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	
	// Emphasis
	Bold   = color.New(color.Bold).SprintFunc()
	Gray   = color.New(color.FgHiBlack).SprintFunc()
	
	// Combined styles
	GreenBold  = color.New(color.FgGreen, color.Bold).SprintFunc()
	CyanBold   = color.New(color.FgCyan, color.Bold).SprintFunc()
	YellowBold = color.New(color.FgYellow, color.Bold).SprintFunc()
	RedBold    = color.New(color.FgRed, color.Bold).SprintFunc()
)
