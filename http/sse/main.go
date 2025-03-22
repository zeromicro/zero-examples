package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type SseHandler struct {
	clients map[chan string]struct{}
}

func NewSseHandler() *SseHandler {
	return &SseHandler{
		clients: make(map[chan string]struct{}),
	}
}

// Serve handles the SSE connection
func (h *SseHandler) Serve(w http.ResponseWriter, r *http.Request) {
	// Set necessary headers for SSE
	// for versions <= v1.8.1, no need for versions > v1.8.1
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	// CORS header is optional here since client and server share origin
	// ctx.SetHeader("Access-Control-Allow-Origin", "*")

	// Create a channel for this client
	clientChan := make(chan string)
	h.clients[clientChan] = struct{}{}

	// Clean up when client disconnects
	defer func() {
		delete(h.clients, clientChan)
		close(clientChan)
	}()

	// Keep the connection alive and send events
	for {
		select {
		case msg := <-clientChan:
			// Send the event to the client
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			// Client disconnected
			return
		}
	}
}

// SimulateEvents generates periodic events
func (h *SseHandler) SimulateEvents() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		message := fmt.Sprintf("Server time: %s", time.Now().Format(time.RFC3339))
		// Broadcast to all connected clients
		for clientChan := range h.clients {
			select {
			case clientChan <- message:
			default:
				// Skip if channel is blocked
			}
		}
	}
}

func main() {
	server := rest.MustNewServer(rest.RestConf{
		Host: "0.0.0.0",
		Port: 8080,
	}, rest.WithFileServer("/static", http.Dir("static")))
	defer server.Stop()

	// Initialize SSE handler
	sseHandler := NewSseHandler()

	// Register SSE endpoint
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/sse",
		Handler: sseHandler.Serve,
	}, rest.WithTimeout(0))

	// for versions > v1.8.1
	// server.AddRoute(rest.Route{
	// 	Method:  http.MethodGet,
	// 	Path:    "/sse",
	// 	Handler: sseHandler.Serve,
	// }, rest.WithSSE())

	// Start event simulator in a separate goroutine
	go sseHandler.SimulateEvents()

	logx.Info("Server starting on :8080")
	server.Start()
}
