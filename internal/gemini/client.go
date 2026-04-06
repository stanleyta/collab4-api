package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	"github.com/alpkeskin/gotoon"
	"radixium.com/collab4-api/internal/contract"
)

// Status constants
const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusPending = "pending"
)

// Client defines the interface for interacting with the Gemini CLI.
type Client interface {
	Execute(ctx context.Context, msg *contract.Message) error
	ToTOON(msg *contract.Message) (string, error)
}

type client struct{}

// NewClient creates a new instance of the Gemini client implementation.
func NewClient() Client {
	return &client{}
}

func (c *client) Execute(ctx context.Context, msg *contract.Message) error {
	// Preparing the command: gemini -p <prompt> -o json
	// We wrap the prompt with context in brackets as an idiomatic pattern.
	args := []string{"-p", fmt.Sprintf("[%s] %s", msg.Context, msg.Prompt), "-o", "json"}
	cmd := exec.CommandContext(ctx, "gemini", args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		msg.Status = StatusError
		msg.Response = stderr.String()
		return fmt.Errorf("gemini command failed: %w, stderr: %s", err, stderr.String())
	}

	msg.Status = StatusSuccess
	msg.Response = out.String()
	return nil
}

func (c *client) ToTOON(msg *contract.Message) (string, error) {
	// Map to interface{} for gotoon.Encode
	data := map[string]interface{}{
		"id":        msg.ID,
		"timestamp": msg.Timestamp.Format(time.RFC3339),
		"context":   msg.Context,
		"prompt":    msg.Prompt,
		"status":    msg.Status,
	}

	// The gemini CLI returns JSON (because of the -o json flag).
	// We unmarshal it here so it gets encoded as native TOON nodes,
	// rather than an escaped JSON string.
	var parsedResponse interface{}
	if err := json.Unmarshal([]byte(msg.Response), &parsedResponse); err == nil {
		data["response"] = parsedResponse
	} else {
		// Fallback to raw string if it's not valid JSON (e.g., during some errors)
		data["response"] = msg.Response
	}

	encoded, err := gotoon.Encode(data)
	if err != nil {
		return "", fmt.Errorf("failed to encode to TOON: %w", err)
	}
	return encoded, nil
}
