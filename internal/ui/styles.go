package ui

import "github.com/charmbracelet/lipgloss"

var (
	primary = lipgloss.Color("#00E5FF")
	accent  = lipgloss.Color("#FF1744")
	gray    = lipgloss.Color("#555555")

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000")).
			Background(primary).
			Bold(true).
			Padding(0, 2).
			MarginBottom(1)

	selectedItem = lipgloss.NewStyle().
			Foreground(primary).
			Bold(true)

	unselectedItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	statusStyle = lipgloss.NewStyle().
			Foreground(accent).
			Bold(true).
			MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			Foreground(gray).
			Italic(true)
)
