package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite" // SQLite driver
)

// Task represents a task in the database
type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

var database *sql.DB

// Init initializes the database connection and creates tables if needed
func Init() error {
	// Get user's home directory for storing the database
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Create .taskman directory if it doesn't exist
	taskmanDir := filepath.Join(homeDir, ".taskman")
	if err := os.MkdirAll(taskmanDir, 0755); err != nil {
		return fmt.Errorf("failed to create taskman directory: %w", err)
	}

	// Database file path
	dbPath := filepath.Join(taskmanDir, "tasks.db")

	// Open database connection
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	database = db

	// Create tables
	if err := createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

// createTables creates the tasks table if it doesn't exist
func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL
	);
	`

	_, err := database.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create tasks table: %w", err)
	}

	return nil
}

// Close closes the database connection
func Close() error {
	if database != nil {
		return database.Close()
	}
	return nil
}

// GetDB returns the database connection
// This will be used by other packages to interact with the database
func GetDB() *sql.DB {
	return database
}
