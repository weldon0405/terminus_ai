package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	APIKey          string
	APIEndpoint     string
	DefaultModel    string
	AvailableModels []string
	MaxTokens       int
}

// Load reads configuration from environment variables or .env file
func Load() (*Config, error) {
	// Try to load .env file, but continue if not found
	_ = godotenv.Load()

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		return nil, errors.New("ANTHROPIC_API_KEY environment variable is required")
	}

	// Get API endpoint with fallback to default
	apiEndpoint := os.Getenv("ANTHROPIC_API_ENDPOINT")
	if apiEndpoint == "" {
		apiEndpoint = "https://api.anthropic.com/v1/messages"
	}

	// Get default model with fallback
	defaultModel := os.Getenv("ANTHROPIC_DEFAULT_MODEL")
	if defaultModel == "" {
		defaultModel = "claude-3-5-sonnet-20240620"
	}

	// Get available models with fallback
	availableModelsStr := os.Getenv("ANTHROPIC_AVAILABLE_MODELS")
	var availableModels []string
	if availableModelsStr == "" {
		availableModels = []string{
			"claude-3-5-sonnet-20240620",
			"claude-3-5-haiku-20240307",
			"claude-3-opus-20240229",
			"claude-3-7-sonnet-20250219",
		}
	} else {
		availableModels = strings.Split(availableModelsStr, ",")
		for i := range availableModels {
			availableModels[i] = strings.TrimSpace(availableModels[i])
		}
	}

	// Check if default model is in available models
	modelFound := false
	for _, model := range availableModels {
		if model == defaultModel {
			modelFound = true
			break
		}
	}

	// If default model not in available models, add it
	if !modelFound && defaultModel != "" {
		availableModels = append(availableModels, defaultModel)
	}

	// Get max tokens with fallback
	maxTokens := 4096 // Default value

	return &Config{
		APIKey:          apiKey,
		APIEndpoint:     apiEndpoint,
		DefaultModel:    defaultModel,
		AvailableModels: availableModels,
		MaxTokens:       maxTokens,
	}, nil
}
