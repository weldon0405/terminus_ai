package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/weldon0405/terminus_ai/internal/api"
	"github.com/weldon0405/terminus_ai/internal/config"
)

// Model represents the application state
type Model struct {
	// Components
	viewport viewport.Model
	textarea textarea.Model
	spinner  spinner.Model

	// State
	messages     []api.Message
	loading      bool
	err          error
	lastResponse *api.Response

	// Window dimensions
	windowHeight int
	windowWidth  int

	// Configuration
	config        *config.Config
	selectedModel string

	// API client
	apiClient *api.Client

	// Styling
	styles Styles
}

// NewModel creates a new Model with the provided configuration
func NewModel(cfg *config.Config) Model {
	// Initialize text area for input
	ta := textarea.New()
	ta.Placeholder = "Ask Claude something..."
	ta.Focus()
	ta.CharLimit = 5000
	ta.SetWidth(30)
	ta.SetHeight(3)

	// Initialize viewport for chat history
	vp := viewport.New(30, 30)
	vp.SetContent("Welcome to Claude TUI! Type your message below and press Ctrl+Enter to send.")

	// Initialize spinner for loading state
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinner.Style.Foreground(spinner.Style.GetForeground())

	// Create API client
	client := api.NewClient(cfg.APIKey, cfg.APIEndpoint)

	return Model{
		textarea:      ta,
		viewport:      vp,
		spinner:       s,
		messages:      []api.Message{},
		loading:       false,
		config:        cfg,
		selectedModel: cfg.DefaultModel,
		apiClient:     client,
		styles:        DefaultStyles(),
	}
}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, m.spinner.Tick)
}
