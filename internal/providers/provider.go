// package providers

// // TODO: Implement providers



package providers

import (
	"context"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Content string `json:"content"`
}

type Provider interface {
	Chat(ctx context.Context, req ChatRequest) (ChatResponse, error)
}