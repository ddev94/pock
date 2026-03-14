package utils

import "github.com/charmbracelet/lipgloss"

var (
	// Success colors
	Green = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Render
	Cyan  = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Render
	Blue  = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render

	// Warning/Error colors
	Yellow = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Render
	Red    = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render

	// Emphasis
	Bold = lipgloss.NewStyle().Bold(true).Render
	Gray = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render

	// Combined styles
	GreenBold  = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true).Render
	CyanBold   = lipgloss.NewStyle().Foreground(lipgloss.Color("6")).Bold(true).Render
	YellowBold = lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Bold(true).Render
	RedBold    = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true).Render
)
