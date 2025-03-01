package api

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request represents the request to Anthropic API
type Request struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

// ContentBlock represents a block of content in the Anthropic API response
type ContentBlock struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// ErrorResponse represents an error returned by the Anthropic API
type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Response represents the response from Anthropic API
type Response struct {
	Content    []ContentBlock `json:"content"`
	Model      string         `json:"model"`
	StopReason string         `json:"stop_reason"`
	Error      ErrorResponse  `json:"error,omitempty"`
}
