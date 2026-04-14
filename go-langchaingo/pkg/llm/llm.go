package llm

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

// ProviderType defines supported LLM backends.
type ProviderType string

const (
	OpenAI    ProviderType = "openai"
	Anthropic ProviderType = "anthropic"
	Gemini    ProviderType = "gemini"
	Ollama    ProviderType = "ollama"
)

// Provider handles LLM operations using langchaingo.
type Provider struct {
	model       llms.Model
	modelName   string
	temperature float64
}

// NewProvider initializes a provider based on the requested type and model name.
func NewProvider(ctx context.Context, pType ProviderType, modelName string, temperature float64) (*Provider, error) {
	var model llms.Model
	var err error
	effectiveModel := modelName

	switch pType {
	case OpenAI:
		opts := []openai.Option{}
		if modelName != "" {
			opts = append(opts, openai.WithModel(modelName))
		} else {
			effectiveModel = "gpt-4o" // Default if not specified
		}
		model, err = openai.New(opts...)
	case Anthropic:
		opts := []anthropic.Option{}
		if modelName != "" {
			opts = append(opts, anthropic.WithModel(modelName))
		} else {
			effectiveModel = "claude-3-5-sonnet-20240620"
		}
		model, err = anthropic.New(opts...)
	case Gemini:
		opts := []googleai.Option{}
		if modelName != "" {
			opts = append(opts, googleai.WithDefaultModel(modelName))
		} else {
			effectiveModel = "gemini-2.5-flash"
		}
		model, err = googleai.New(ctx, opts...)
	case Ollama:
		opts := []ollama.Option{}
		if modelName != "" {
			opts = append(opts, ollama.WithModel(modelName))
		} else {
			effectiveModel = "llama3"
			opts = append(opts, ollama.WithModel(effectiveModel))
		}
		model, err = ollama.New(opts...)
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", pType)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create %s model (%s): %w", pType, effectiveModel, err)
	}

	return &Provider{model: model, modelName: effectiveModel, temperature: temperature}, nil
}

// ModelName returns the name of the model being used.
func (p *Provider) ModelName() string {
	return p.modelName
}

// Generate sends a prompt to the model and returns the response.
func (p *Provider) Generate(ctx context.Context, prompt string) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, p.model, prompt,
		llms.WithMaxTokens(1024),
		llms.WithTemperature(p.temperature),
	)
}

// MinimalTest just prints something to show it's alive.
func (p *Provider) MinimalTest(ctx context.Context) string {
	return "LLM Provider initialized successfully with langchaingo"
}
