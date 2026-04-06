package main

import (
	"testing"
	"time"

	"radixium.com/collab4-api/internal/contract"
)

func TestMessageContract(t *testing.T) {
	msg := &contract.Message{
		ID:        "test-1",
		Timestamp: time.Now(),
		Context:   "Testing",
		Prompt:    "Verify the contract",
		Status:    "pending",
	}

	if msg.ID != "test-1" {
		t.Errorf("Expected ID test-1, got %s", msg.ID)
	}
}
