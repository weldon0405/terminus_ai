package tui

import (
	"fmt"
	"strings"
)

// View renders the user interface
func (m Model) View() string {
	var b strings.Builder

	// Application title
	title := m.styles.Title.Width(m.windowWidth).Render("Claude TUI")
	b.WriteString(title)

	// Info line with selected model and keybindings
	info := m.styles.Info.Render(fmt.Sprintf(
		"Model: %s (Press Tab to change) | Enter to send | Ctrl+C to quit",
		m.selectedModel,
	))
	b.WriteString("\n" + info + "\n\n")

	// Chat viewport
	b.WriteString(m.viewport.View())
	b.WriteString("\n\n")

	// Input area with loading indicator when applicable
	inputPrompt := m.styles.InputPrompt.Render("Your message:")
	if m.loading {
		loadingText := m.styles.LoadingText.Render("Thinking...")
		inputPrompt = fmt.Sprintf("%s %s", m.spinner.View(), loadingText)
	}
	b.WriteString(inputPrompt + "\n")
	b.WriteString(m.textarea.View())

	// Error message if any
	if m.err != nil {
		b.WriteString("\n" + m.styles.Error.Render(m.err.Error()))
	}

	return b.String()
}
