package db

import (
	"os"
	"testing"
)

// setupTestDB creates a temporary database for testing
func setupTestDB(t *testing.T) func() {
	// Create temporary directory for test database
	tempDir := t.TempDir()

	// Set up environment to use test database
	oldHomeDir := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)

	// Initialize database
	if err := Init(); err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	// Return cleanup function
	return func() {
		Close()
		os.Setenv("HOME", oldHomeDir)
	}
}

func TestAddTask(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	tests := []struct {
		name        string
		title       string
		description string
		wantErr     bool
	}{
		{
			name:        "Add task with title only",
			title:       "Test Task",
			description: "",
			wantErr:     false,
		},
		{
			name:        "Add task with title and description",
			title:       "Test Task 2",
			description: "Test Description",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := AddTask(tt.title, tt.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if task.Title != tt.title {
					t.Errorf("AddTask() title = %v, want %v", task.Title, tt.title)
				}
				if task.Description != tt.description {
					t.Errorf("AddTask() description = %v, want %v", task.Description, tt.description)
				}
				if task.Completed {
					t.Errorf("AddTask() completed = %v, want false", task.Completed)
				}
			}
		})
	}
}

func TestGetAllTasks(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// Add some test tasks
	AddTask("Task 1", "Description 1")
	AddTask("Task 2", "Description 2")
	AddTask("Task 3", "")

	tasks, err := GetAllTasks()
	if err != nil {
		t.Fatalf("GetAllTasks() error = %v", err)
	}

	if len(tasks) != 3 {
		t.Errorf("GetAllTasks() returned %d tasks, want 3", len(tasks))
	}

	// Verify tasks are ordered by created_at DESC (newest first)
	if tasks[0].Title != "Task 3" {
		t.Errorf("GetAllTasks() first task title = %v, want Task 3", tasks[0].Title)
	}
}

func TestCompleteTask(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// Add a test task
	task, _ := AddTask("Test Task", "")

	// Complete the task
	err := CompleteTask(task.ID)
	if err != nil {
		t.Errorf("CompleteTask() error = %v", err)
	}

	// Verify task is completed
	tasks, _ := GetAllTasks()
	if len(tasks) > 0 && !tasks[0].Completed {
		t.Errorf("CompleteTask() did not mark task as completed")
	}

	// Try to complete non-existent task
	err = CompleteTask(9999)
	if err == nil {
		t.Errorf("CompleteTask() should return error for non-existent task")
	}
}

func TestDeleteTask(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	// Add a test task
	task, _ := AddTask("Test Task", "")

	// Delete the task
	err := DeleteTask(task.ID)
	if err != nil {
		t.Errorf("DeleteTask() error = %v", err)
	}

	// Verify task is deleted
	tasks, _ := GetAllTasks()
	if len(tasks) != 0 {
		t.Errorf("DeleteTask() did not delete task, found %d tasks", len(tasks))
	}

	// Try to delete non-existent task
	err = DeleteTask(9999)
	if err == nil {
		t.Errorf("DeleteTask() should return error for non-existent task")
	}
}
