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
	providerFlag := flag.String("provider", string(llm.Ollama), "LLM provider to use (openai, anthropic, gemini, ollama)")
	modelFlag := flag.String("model", "llama3", "Specific model name (e.g., gemini-2.5-flash, claude-3-5-sonnet-20240620, llama3)")
	promptFlag := flag.String("prompt", "Tell me a short joke.", "The prompt to send to the LLM")
	tempFlag := flag.Float64("temp", 1.2, "Temperature for the LLM (0.0 to 2.0)")
	flag.Parse()

	ctx := context.Background()

	// Initialize the service
	svc, err := service.NewService(ctx, llm.ProviderType(*providerFlag), *modelFlag, *tempFlag)
	if err != nil {
		log.Printf("Warning: Service initialization failed: %v", err)
		os.Exit(1)
	}

	log.Printf("Service initialized. Provider: %s, Model: %s, Temperature: %.2f", *providerFlag, svc.ModelName(), *tempFlag)


	if err := svc.Run(ctx); err != nil {
		log.Fatalf("Service execution failed: %v", err)
	}

	// Example: Process a prompt
	response, err := svc.Process(ctx, *promptFlag)
	if err != nil {
		log.Printf("Process failed (expected if no API key or invalid model): %v", err)
	} else {
		log.Printf("Model responded: %s", response)
	}
}
