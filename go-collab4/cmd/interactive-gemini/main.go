package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"radixium.com/go-collab4/pkg/contract"
	"radixium.com/go-collab4/pkg/gemini"
)

func main() {
	client := gemini.NewClient()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Interactive Gemini CLI")
	fmt.Println("Type your prompt and press Enter. Type 'exit' to quit.")

	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}

		if strings.TrimSpace(input) == "" {
			continue
		}

		msg := &contract.Message{
			ID:        fmt.Sprintf("interactive-%d", time.Now().UnixNano()),
			Timestamp: time.Now(),
			Context:   "Interactive Session",
			Prompt:    input,
			Status:    gemini.StatusPending,
		}

		fmt.Println("Processing...")

		// Use a reasonable timeout for each interaction
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		err := client.Execute(ctx, msg)
		cancel()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			// Continue to show what we have in TOON even on error
		}

		toonOutput, err := client.ToTOON(msg)
		if err != nil {
			log.Printf("Error encoding to TOON: %v", err)
			continue
		}

		fmt.Println("\n--- Structured Message (TOON) ---")
		fmt.Println(toonOutput)
	}

	fmt.Println("Goodbye!")
}
