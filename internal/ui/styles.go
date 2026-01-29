package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	primary   = lipgloss.Color("#00E5FF")
	secondary = lipgloss.Color("#00B8D4")
	success   = lipgloss.Color("#00E676")
	errorClr  = lipgloss.Color("#FF1744")
	subtle    = lipgloss.Color("#555555")
	white     = lipgloss.Color("#FFFFFF")

	// Container with rounded border
	containerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondary).
			Padding(1, 2)

	// Title
	titleStyle = lipgloss.NewStyle().
			Foreground(primary).
			Bold(true).
			MarginBottom(1)

	// Section header
	sectionStyle = lipgloss.NewStyle().
			Foreground(white).
			Bold(true).
			MarginBottom(1)

	// Status messages
	successStyle = lipgloss.NewStyle().
			Foreground(success).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorClr).
			Bold(true)

	// Active badge
	activeBadgeStyle = lipgloss.NewStyle().
				Foreground(success)

	// Help text
	helpStyle = lipgloss.NewStyle().
			Foreground(subtle).
			Italic(true).
			MarginTop(1)

	// List item styles
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(primary).
				Bold(true)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(white)
)
