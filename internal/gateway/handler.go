// package gateway

// // TODO: Implement gateway





package gateway

import (
	"encoding/json"
	"net/http"

	"github.com/Shyam525/nexus-gateway/internal/providers"
)

type Handler struct {
	provider providers.Provider
}

func NewHandler(provider providers.Provider) *Handler {
	return &Handler{
		provider: provider,
	}
}

func (h *Handler) Chat(w http.ResponseWriter, r *http.Request) {
	var req providers.ChatRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.provider.Chat(r.Context(), req)
	if err != nil {
		http.Error(w, "provider error", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}