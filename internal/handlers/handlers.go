package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/oodles-noodles/secure-supply-chain-go/pkg/logger"
	"gopkg.in/yaml.v3"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status" yaml:"status"`
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
	Version   string    `json:"version" yaml:"version"`
}

// HealthHandler handles health check requests
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// InfoResponse represents application information
type InfoResponse struct {
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Version     string            `json:"version" yaml:"version"`
	Features    []string          `json:"features" yaml:"features"`
	Metadata    map[string]string `json:"metadata" yaml:"metadata"`
}

// InfoHandler provides application information
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		Name:        "Secure Supply Chain Demo",
		Description: "A demonstration of secure software supply chain practices",
		Version:     "1.0.0",
		Features: []string{
			"RESTful API",
			"Structured Logging",
			"Prometheus Metrics",
			"Configuration Management",
			"Database Integration",
			"JWT Authentication Support",
		},
		Metadata: map[string]string{
			"go_version": "1.24.12",
			"build_date": time.Now().Format(time.RFC3339),
		},
	}

	// Support both JSON and YAML output based on Accept header
	accept := r.Header.Get("Accept")
	if accept == "application/yaml" || accept == "text/yaml" {
		w.Header().Set("Content-Type", "application/yaml")
		yaml.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// User represents a user in the system
type User struct {
	ID        string    `json:"id" yaml:"id"`
	Username  string    `json:"username" yaml:"username"`
	Email     string    `json:"email" yaml:"email"`
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}

// UsersHandler handles user-related requests
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("Fetching users list")

	// Mock data for demonstration
	users := []User{
		{
			ID:        uuid.New().String(),
			Username:  "admin",
			Email:     "admin@example.com",
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        uuid.New().String(),
			Username:  "user1",
			Email:     "user1@example.com",
			CreatedAt: time.Now().Add(-12 * time.Hour),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// WelcomeHandler handles the welcome/root endpoint
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Welcome to the Secure Supply Chain Demo API",
		"version": "1.0.0",
		"endpoints": []string{
			"/health",
			"/info",
			"/api/users",
			"/metrics",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
