package service

import (
	"context"
	"fmt"

	"radixium.com/go-langchaingo/pkg/llm"
)

// Service represents the application logic.
type Service struct {
	llmProvider *llm.Provider
}

// NewService creates a new instance of Service with a specific LLM provider and model.
func NewService(ctx context.Context, pType llm.ProviderType, modelName string, temperature float64) (*Service, error) {
	provider, err := llm.NewProvider(ctx, pType, modelName, temperature)
	if err != nil {
		return nil, fmt.Errorf("could not initialize llm provider (%s/%s): %w", pType, modelName, err)
	}
	return &Service{llmProvider: provider}, nil
}

// ModelName returns the model name being used by the service.
func (s *Service) ModelName() string {
	return s.llmProvider.ModelName()
}

// Run starts the service logic.
func (s *Service) Run(ctx context.Context) error {
	fmt.Println("Service is running...")
	msg := s.llmProvider.MinimalTest(ctx)
	fmt.Println(msg)
	return nil
}

// Process handles a single prompt request.
func (s *Service) Process(ctx context.Context, prompt string) (string, error) {
	fmt.Printf("Processing prompt: %s\n", prompt)
	return s.llmProvider.Generate(ctx, prompt)
}
