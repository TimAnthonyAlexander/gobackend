package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yourusername/gobackend/internal/config"
	"github.com/yourusername/gobackend/pkg/utils"
)

func main() {
	// Initialize loggers
	utils.InitLoggers()

	// Load configuration
	cfg := config.NewConfig()

	// Create and start the server
	server := NewServer(cfg)
	if err := server.Start(); err != nil {
		utils.ErrorLogger.Fatalf("Failed to start server: %v", err)
	}

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		utils.ErrorLogger.Fatalf("Server forced to shutdown: %v", err)
	}

	utils.InfoLogger.Println("Server exited properly")
} 