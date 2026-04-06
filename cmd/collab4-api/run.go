package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"radixium.com/collab4-api/internal/contract"
	"radixium.com/collab4-api/internal/gemini"
)

func main() {
	var (
		prompt     = flag.String("p", "Check if this code follows Go idioms.", "The prompt to send to Gemini.")
		contextStr = flag.String("c", "Skill: Code Review", "The context for the prompt.")
		timeout    = flag.Duration("t", 30*time.Second, "Timeout for the gemini command.")
	)
	flag.Parse()

	client := gemini.NewClient()

	msg := &contract.Message{
		ID:        fmt.Sprintf("run-%d", time.Now().UnixNano()),
		Timestamp: time.Now(),
		Context:   *contextStr,
		Prompt:    *prompt,
		Status:    "pending",
	}

	fmt.Printf("Executing instruction: %s\n", msg.Prompt)

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	if err := client.Execute(ctx, msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}

	// Structured output for observability/logging in TOON
	toonOutput, err := client.ToTOON(msg)
	if err != nil {
		log.Fatalf("Fatal error encoding to TOON: %v", err)
	}

	fmt.Println("\n--- Structured Message (TOON) ---")
	fmt.Println(toonOutput)

	if msg.Status == gemini.StatusError {
		os.Exit(1)
	}
}
