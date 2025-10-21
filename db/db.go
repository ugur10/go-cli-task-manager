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

// AddTask creates a new task in the database
func AddTask(title, description string) (*Task, error) {
	query := `INSERT INTO tasks (title, description, completed, created_at)
	          VALUES (?, ?, 0, ?)`

	now := time.Now()
	result, err := database.Exec(query, title, description, now)
	if err != nil {
		return nil, fmt.Errorf("failed to insert task: %w", err)
	}

	// Get the ID of the newly created task
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	task := &Task{
		ID:          int(id),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   now,
	}

	return task, nil
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks() ([]Task, error) {
	query := `SELECT id, title, description, completed, created_at
	          FROM tasks
	          ORDER BY created_at DESC`

	rows, err := database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tasks: %w", err)
	}

	return tasks, nil
}

// CompleteTask marks a task as completed by its ID
func CompleteTask(id int) error {
	query := `UPDATE tasks SET completed = 1 WHERE id = ?`

	result, err := database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}

	// Check if any row was actually updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return nil
}

// DeleteTask removes a task from the database by its ID
func DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = ?`

	result, err := database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	// Check if any row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return nil
}
