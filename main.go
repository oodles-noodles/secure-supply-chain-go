package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/oodles-noodles/secure-supply-chain-go/internal/config"
	"github.com/oodles-noodles/secure-supply-chain-go/internal/database"
	"github.com/oodles-noodles/secure-supply-chain-go/internal/handlers"
	"github.com/oodles-noodles/secure-supply-chain-go/internal/middleware"
	"github.com/oodles-noodles/secure-supply-chain-go/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	logger.Init(cfg.Logging.Level)
	log := logger.GetLogger()

	log.Info("Starting Secure Supply Chain Demo Application")
	log.Infof("Server will listen on %s:%s", cfg.Server.Host, cfg.Server.Port)

	// Initialize database connection (non-fatal if it fails)
	if err := database.Connect(cfg.Database); err != nil {
		log.Warnf("Database connection failed: %v", err)
	}
	defer database.Close()

	// Setup router
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.LoggingMiddleware)

	// Register routes
	r.HandleFunc("/", handlers.WelcomeHandler).Methods("GET")
	r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	r.HandleFunc("/info", handlers.InfoHandler).Methods("GET")
	r.HandleFunc("/api/users", handlers.UsersHandler).Methods("GET")
	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	// Create HTTP server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Infof("Server starting on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Info("Server started successfully")

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Info("Server stopped")
}
