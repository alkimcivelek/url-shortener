package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"url-shortener/internal/adapter/http/handler"
	"url-shortener/internal/adapter/http/router"
	"url-shortener/internal/application/service"
	"url-shortener/internal/infrastructure/config"
	"url-shortener/internal/infrastructure/persistence/memory"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize dependencies
	urlRepository := memory.NewURLRepositoryMemory()
	urlService := service.NewURLApplicationService(urlRepository)
	urlHandler := handler.NewURLHandler(urlService, cfg.BaseURL)
	httpRouter := router.NewRouter(urlHandler)

	// Create HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      httpRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	go func() {
		log.Printf("Starting service on port %s", cfg.Port)
		log.Printf("Health check: %s/health", cfg.BaseURL)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
