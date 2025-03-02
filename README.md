# Terminus AI

A Terminal User Interface (TUI) application for interacting with Anthropic's Claude LLM API directly from your terminal.

## Features

- Clean, intuitive terminal interface
- Real-time conversation with Claude
- Support for multiple Claude models (Claude 3.5 Sonnet, Claude 3.5 Haiku, Claude 3 Opus, Claude 3.7 Sonnet)
- Model switching with Tab key
- Rich text formatting for improved readability
- Error handling and API request management

## Screenshots

[Screenshots would be placed here]

## Prerequisites

- Go 1.18 or higher
- An Anthropic API key

## Installation

1. Clone this repository:

   ```
   git clone https://github.com/weldon0405/terminus_ai.git
   cd claude-tui
   ```

2. Install dependencies:

   ```
   go mod download
   ```

3. Set up your API key:

   Either create a `.env` file in the project root with:

   ```
   ANTHROPIC_API_KEY=your_api_key_here
   ```

   Or set it as an environment variable:

   ```
   export ANTHROPIC_API_KEY=your_api_key_here
   ```

4. Build the application:
   ```
   go build ./cmd/terminus_ai
   ```

## Usage

Run the application:

```
./terminus_ai
```

### Keyboard Controls

- **Type** to compose your message
- **Enter** to send message to Claude
- **Tab** to cycle through available Claude models
- **Ctrl+C** or **Esc** to quit the application

## Configuration

You can modify the following in the code:

- Available models
- API endpoint
- Default model
- UI styling and colors

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [godotenv](https://github.com/joho/godotenv) - Environment variable loading

## Future Improvements

- Chat history persistence
- Configuration file for customizing appearance
- System prompt customization
- Message editing capabilities
- Response streaming support
- File uploads for context
- Better handling of long responses and code blocks

## License

MIT

## Acknowledgements

- [Anthropic](https://www.anthropic.com/) for the Claude API
- [Charm](https://charm.sh/) for the excellent TUI libraries
