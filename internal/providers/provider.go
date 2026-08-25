package providers

import (
	"context"
	"fmt"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

func (r ChatRequest) Validate() error {
	if r.Model == "" {
		return fmt.Errorf("model is required")
	}

	if len(r.Messages) == 0 {
		return fmt.Errorf("messages cannot be empty")
	}

	for _, message := range r.Messages {
		if message.Role == "" {
			return fmt.Errorf("message role is required")
		}

		if message.Role != "system" &&
			message.Role != "user" &&
			message.Role != "assistant" {
			return fmt.Errorf("invalid message role: %s", message.Role)
		}

		if message.Content == "" {
			return fmt.Errorf("message content is required")
		}
	}

	return nil
}

type ChatResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Content string `json:"content"`
}

type Provider interface {
	Chat(ctx context.Context, req ChatRequest) (ChatResponse, error)
}