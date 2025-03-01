package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/weldon0405/terminus_ai/internal/api"
)

// ChatModel represents the chat history component
type ChatModel struct {
	viewport       viewport.Model
	messages       []api.Message
	userStyle      lipgloss.Style
	assistantStyle lipgloss.Style
	errorStyle     lipgloss.Style
	width          int
	height         int
}

// NewChatModel creates a new chat history component
func NewChatModel(width, height int) ChatModel {
	vp := viewport.New(width, height)
	vp.SetContent("Welcome to Claude TUI! Type your message below and press Ctrl+Enter to send.")

	return ChatModel{
		viewport: vp,
		messages: []api.Message{},
		userStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5F9EA0")).
			Bold(true),
		assistantStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9370DB")),
		errorStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")),
		width:  width,
		height: height,
	}
}

// Init initializes the component
func (m ChatModel) Init() tea.Cmd {
	return nil
}

// Update handles events for the component
func (m ChatModel) Update(msg tea.Msg) (ChatModel, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)

	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = m.width
		m.viewport.Height = m.height
	}

	return m, cmd
}

// View renders the component
func (m ChatModel) View() string {
	return m.viewport.View()
}

// AddUserMessage adds a user message to the chat
func (m *ChatModel) AddUserMessage(content string) {
	// Create and add the message
	message := api.Message{
		Role:    "user",
		Content: content,
	}
	m.messages = append(m.messages, message)

	// Update the viewport content
	vpContent := m.viewport.Content
	vpContent += "\n\n" + m.userStyle.Render("You: ") + "\n" + content
	m.viewport.SetContent(vpContent)
	m.viewport.GotoBottom()
}

// AddAssistantMessage adds an assistant message to the chat
func (m *ChatModel) AddAssistantMessage(content string) {
	// Create and add the message
	message := api.Message{
		Role:    "assistant",
		Content: content,
	}
	m.messages = append(m.messages, message)

	// Update the viewport content
	vpContent := m.viewport.Content
	vpContent += "\n\n" + m.assistantStyle.Render("Claude: ") + "\n" + content
	m.viewport.SetContent(vpContent)
	m.viewport.GotoBottom()
}

// AddErrorMessage adds an error message to the chat
func (m *ChatModel) AddErrorMessage(err error) {
	// Update the viewport content
	vpContent := m.viewport.Content
	vpContent += "\n\n" + m.errorStyle.Render(fmt.Sprintf("Error: %v", err))
	m.viewport.SetContent(vpContent)
	m.viewport.GotoBottom()
}

// GetMessages returns all messages in the chat
func (m *ChatModel) GetMessages() []api.Message {
	return m.messages
}

// SetSize sets the size of the component
func (m *ChatModel) SetSize(width, height int) {
	m.width = width
	m.height = height
	m.viewport.Width = width
	m.viewport.Height = height
}

// Resize adjusts the component size
func (m *ChatModel) Resize(width, height int) {
	m.SetSize(width, height)

	// Force re-render with new dimensions
	content := m.viewport.Content
	m.viewport.SetContent(strings.TrimSpace(content))
}
