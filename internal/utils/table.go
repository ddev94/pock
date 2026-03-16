package utils

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"golang.org/x/term"
)

// collapseNewlines replaces newlines/tabs with a single space for single-line display.
func collapseNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", "  ")
	return s
}

// getTermWidth returns the terminal width or a default of 80.
func getTermWidth() int {
	fd := int(uintptr(0)) // stdin
	if w, _, err := term.GetSize(fd); err == nil && w > 0 {
		return w
	}
	return 80
}

// RenderTable renders data as a clean table using lipgloss.
// Cells are automatically truncated with "…" to fit terminal width.
func RenderTable(headers []string, rows [][]string) {
	termWidth := getTermWidth()
	const gap = 2 // Space between columns

	// Collapse newlines and strip ANSI from all cells.
	normalizedRows := make([][]string, len(rows))
	for i, row := range rows {
		normalizedRows[i] = make([]string, len(row))
		for j, cell := range row {
			normalizedRows[i][j] = collapseNewlines(ansi.Strip(cell))
		}
	}

	// Calculate initial column widths based on content.
	widths := make([]int, len(headers))
	idealWidths := make(map[string]int)
	idealWidths["Name"] = 20
	idealWidths["Command"] = 50
	idealWidths["Description"] = 30

	for i, h := range headers {
		widths[i] = len(h)
		if ideal, ok := idealWidths[h]; ok && widths[i] < ideal {
			widths[i] = ideal
		}
	}
	for _, row := range normalizedRows {
		for j, cell := range row {
			if j < len(widths) {
				if w := len(cell); w > widths[j] {
					widths[j] = w
				}
			}
		}
	}

	// Cap maximum widths to prevent excessive expansion.
	maxWidths := make(map[string]int)
	maxWidths["Name"] = 30
	maxWidths["Command"] = 80
	maxWidths["Description"] = 50

	for i, h := range headers {
		if max, ok := maxWidths[h]; ok && widths[i] > max {
			widths[i] = max
		}
	}

	// Calculate total width needed and adjust if exceeds terminal.
	totalWidth := 0
	for _, w := range widths {
		totalWidth += w
	}
	totalWidth += gap * (len(headers) - 1)

	// If exceeds terminal, shrink columns proportionally.
	if totalWidth > termWidth-4 {
		available := termWidth - 4 - gap*(len(headers)-1)
		if available < len(headers)*6 {
			// Terminal too narrow, use absolute minimum per column.
			for i := range widths {
				widths[i] = 6
			}
		} else {
			// Proportional shrink while maintaining readability.
			for i := range widths {
				newWidth := (widths[i] * available) / totalWidth
				if newWidth < 8 {
					newWidth = 8
				}
				widths[i] = newWidth
			}

			// Final check - if still too wide, force fit.
			totalWidth = 0
			for _, w := range widths {
				totalWidth += w
			}
			totalWidth += gap * (len(headers) - 1)

			if totalWidth > termWidth-4 {
				for i := range widths {
					widths[i] = (widths[i] * available) / totalWidth
					if widths[i] < 6 {
						widths[i] = 6
					}
				}
			}
		}
	}

	headerStyle := lipgloss.NewStyle().Bold(true)
	sepStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	// Render header row - truncate then pad.
	var headerParts []string
	for i, h := range headers {
		truncated := ansi.Truncate(h, widths[i], "…")
		// Pad to exact width using spaces manually.
		visibleLen := len(truncated)
		padding := 0
		if visibleLen < widths[i] {
			padding = widths[i] - visibleLen
		}
		padded := truncated + strings.Repeat(" ", padding)
		headerParts = append(headerParts, headerStyle.Render(padded))
	}
	fmt.Println(strings.Join(headerParts, strings.Repeat(" ", gap)))

	// Render separator line.
	var sepParts []string
	for _, w := range widths {
		sepParts = append(sepParts, strings.Repeat("─", w))
	}
	fmt.Println(sepStyle.Render(strings.Join(sepParts, strings.Repeat("─", gap))))

	// Render data rows with colors applied after truncation.
	for _, row := range normalizedRows {
		var cellParts []string
		for j := range headers {
			cell := ""
			if j < len(row) {
				cell = row[j]
			}
			truncated := ansi.Truncate(cell, widths[j], "…")

			// Apply colors after truncation.
			switch headers[j] {
			case "Name":
				truncated = Green(truncated)
			case "Command":
				truncated = Yellow(truncated)
			case "Status":
				if strings.EqualFold(strings.TrimSpace(cell), "failure") {
					truncated = Red(truncated)
				} else {
					truncated = Green(truncated)
				}
			}

			// Pad to exact width (accounting for ANSI codes).
			visibleLen := lipgloss.Width(truncated)
			if visibleLen < widths[j] {
				padded := truncated + strings.Repeat(" ", widths[j]-visibleLen)
				cellParts = append(cellParts, padded)
			} else {
				cellParts = append(cellParts, truncated)
			}
		}
		fmt.Println(strings.Join(cellParts, strings.Repeat(" ", gap)))
	}
}
