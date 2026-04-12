package main

import (
	"context"
	"flag"
	"log"
	"os"

	"radixium.com/go-langchaingo/pkg/llm"
	"radixium.com/go-langchaingo/service"
)

func main() {
	// Define flags
	providerFlag := flag.String("provider", string(llm.Gemini), "LLM provider to use (openai, anthropic, gemini)")
	modelFlag := flag.String("model", "gemini-2.5-flash", "Specific model name (e.g., gemini-2.5-flash, claude-3-5-sonnet-20240620)")
	flag.Parse()

	ctx := context.Background()

	log.Printf("Starting service with provider: %s, model: %s", *providerFlag, *modelFlag)

	// Initialize the service
	svc, err := service.NewService(ctx, llm.ProviderType(*providerFlag), *modelFlag)
	if err != nil {
		log.Printf("Warning: Service initialization failed: %v", err)
		os.Exit(1)
	}

	if err := svc.Run(ctx); err != nil {
		log.Fatalf("Service execution failed: %v", err)
	}

	// Example: Process a prompt
	response, err := svc.Process(ctx, "Hello, how are you?")
	if err != nil {
		log.Printf("Process failed (expected if no API key or invalid model): %v", err)
	} else {
		log.Printf("Model responded: %s", response)
	}
}
