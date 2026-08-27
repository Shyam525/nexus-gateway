package gateway

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat", handler.Chat)
	return mux
}
