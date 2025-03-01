package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles contains all the styling for the TUI
type Styles struct {
	Title        lipgloss.Style
	Info         lipgloss.Style
	Error        lipgloss.Style
	UserMsg      lipgloss.Style
	AssistantMsg lipgloss.Style
	InputPrompt  lipgloss.Style
	LoadingText  lipgloss.Style
	ModelSelect  lipgloss.Style
}

// DefaultStyles returns the default styling for the application
func DefaultStyles() Styles {
	return Styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#8A2BE2")).
			Padding(0, 1).
			MarginBottom(1),

		Info: lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color("#666666")),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")),

		UserMsg: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5F9EA0")).
			Bold(true),

		AssistantMsg: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9370DB")),

		InputPrompt: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")),

		LoadingText: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Italic(true),

		ModelSelect: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00CED1")),
	}
}
