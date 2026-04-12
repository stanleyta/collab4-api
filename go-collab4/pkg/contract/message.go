package contract

import (
	"time"
)

// Message represents the structured contract for communication with Gemini.
// We use TOON tags for efficient LLM context serialization.
type Message struct {
	ID        string      `json:"id" toon:"id"`
	Timestamp time.Time   `json:"timestamp" toon:"timestamp"`
	Context   string      `json:"context" toon:"context"`
	Prompt    string      `json:"prompt" toon:"prompt"`
	Response  interface{} `json:"response" toon:"response"`
	Status    string      `json:"status" toon:"status"`
}
