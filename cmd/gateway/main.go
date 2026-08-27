package main

import (
	"log"
	"net/http"

	"github.com/Shyam525/nexus-gateway/internal/gateway"
)

func main() {
	handler := gateway.NewHandler()
	mux := http.NewServeMux()
	mux.Handle("/v1/", gateway.NewRouter(handler))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("NEXUS gateway listening on :8081")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
