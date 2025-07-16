package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/yourusername/gobackend/internal/config"
	"github.com/yourusername/gobackend/internal/handlers"
	"github.com/yourusername/gobackend/pkg/utils"
)

// Server represents the API server
type Server struct {
	config     *config.Config
	httpServer *http.Server
}

// NewServer creates a new API server
func NewServer(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

// setupRoutes configures the API routes
func (s *Server) setupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/datetime", handlers.DateTimeHandler)
	return mux
}

// Start starts the API server
func (s *Server) Start() error {
	router := s.setupRoutes()
	
	addr := fmt.Sprintf("%s:%s", s.config.Server.Host, s.config.Server.Port)
	s.httpServer = &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	utils.InfoLogger.Printf("Starting server on %s", addr)
	
	// Start the server in a goroutine
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.ErrorLogger.Fatalf("Error starting server: %v", err)
		}
	}()

	return nil
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	utils.InfoLogger.Println("Server is shutting down...")
	return s.httpServer.Shutdown(ctx)
} 