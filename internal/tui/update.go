package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/weldon0405/terminus_ai/internal/api"
)

// API response message
type apiResponseMsg struct {
	response *api.Response
	err      error
}

// Command to send a message to the API
func (m Model) sendMessageCmd() tea.Cmd {
	return func() tea.Msg {
		// Send the message
		response, err := m.apiClient.SendMessage(
			m.selectedModel,
			m.messages,
			m.config.MaxTokens,
		)

		// Return the response or error
		return apiResponseMsg{
			response: response,
			err:      err,
		}
	}
}

// Update handles events and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		taCmd tea.Cmd
		vpCmd tea.Cmd
		spCmd tea.Cmd
		cmds  []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			// For Enter key, check if user input is not empty
			if !m.loading && strings.TrimSpace(m.textarea.Value()) != "" {
				if !m.loading && strings.TrimSpace(m.textarea.Value()) != "" {
					// Create user message
					userMsg := api.Message{
						Role:    "user",
						Content: m.textarea.Value(),
					}

					// Add to message history
					m.messages = append(m.messages, userMsg)

					// Update viewport content
					vpContent := m.viewport.View()
					vpContent += "\n\n" + m.styles.UserMsg.Render("You: ") + "\n" + m.textarea.Value()
					m.viewport.SetContent(vpContent)
					m.viewport.GotoBottom()

					// Clear the textarea and start loading
					m.textarea.Reset()
					m.loading = true

					// Send the message to the API
					return m, m.sendMessageCmd()
				}
			}

		case tea.KeyTab:
			// Cycle through available models
			for i, model := range m.config.AvailableModels {
				if model == m.selectedModel {
					nextIndex := (i + 1) % len(m.config.AvailableModels)
					m.selectedModel = m.config.AvailableModels[nextIndex]
					break
				}
			}
		}

	case apiResponseMsg:
		m.loading = false
		if msg.err != nil {
			// Handle error
			m.err = msg.err
			errorMsg := m.styles.Error.Render(fmt.Sprintf("Error: %v", msg.err))
			newContent := m.viewport.View() + "\n\n" + errorMsg
			m.viewport.SetContent(newContent)
			m.viewport.GotoBottom()
		} else {
			// Handle successful response
			m.lastResponse = msg.response

			if len(msg.response.Content) > 0 {
				// Create assistant message
				assistantMsg := api.Message{
					Role:    "assistant",
					Content: msg.response.Content[0].Text,
				}

				// Add to message history
				m.messages = append(m.messages, assistantMsg)

				// Update viewport
				vpContent := m.viewport.View()
				vpContent += "\n\n" + m.styles.AssistantMsg.Render("Claude: ") + "\n" + msg.response.Content[0].Text
				m.viewport.SetContent(vpContent)
				m.viewport.GotoBottom()
			} else if msg.response.Error.Message != "" {
				// Handle API error
				errorMsg := m.styles.Error.Render(fmt.Sprintf("API Error: %s", msg.response.Error.Message))
				newContent := m.viewport.View() + "\n\n" + errorMsg
				m.viewport.SetContent(newContent)
				m.viewport.GotoBottom()
			}
		}

	case tea.WindowSizeMsg:
		// Update window dimensions
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width

		// Adjust viewport and textarea dimensions
		headerHeight := 6 // Title + info
		inputHeight := 5  // Textarea + padding
		footerHeight := 2 // Status line

		contentHeight := m.windowHeight - headerHeight - inputHeight - footerHeight
		contentWidth := m.windowWidth - 4 // Account for margins

		if contentHeight > 0 {
			m.viewport.Height = contentHeight
			m.viewport.Width = contentWidth
		}

		if m.windowWidth > 30 {
			m.textarea.SetWidth(contentWidth)
		}

		// Re-render the viewport with the right dimensions
		currentContent := m.viewport.View()
		m.viewport.SetContent(currentContent)

	case spinner.TickMsg:
		if m.loading {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}

	// Handle textarea and viewport updates
	m.textarea, taCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	if m.loading {
		m.spinner, spCmd = m.spinner.Update(msg)
		cmds = append(cmds, spCmd)
	}

	cmds = append(cmds, taCmd, vpCmd)
	return m, tea.Batch(cmds...)
}
