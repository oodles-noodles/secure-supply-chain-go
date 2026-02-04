package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/oodles-noodles/secure-supply-chain-go/internal/config"
	"github.com/oodles-noodles/secure-supply-chain-go/pkg/logger"
)

// DB is the global database connection
var DB *sql.DB

// Connect establishes a connection to the PostgreSQL database
func Connect(cfg config.DatabaseConfig) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Ping the database to verify the connection (but don't fail if unreachable)
	if err := DB.Ping(); err != nil {
		logger.GetLogger().Warnf("Database ping failed: %v (continuing anyway)", err)
		return nil
	}

	logger.GetLogger().Info("Database connection established")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return DB
}
