package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ugur10/go-cli-task-manager/db"
)

var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "A simple CLI task manager",
	Long: `TaskMan is a command-line task management application built with Go.
It allows you to add, list, complete, and delete tasks with SQLite persistence.`,
}

var (
	// Flag for add command
	taskDescription string
)

var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Long:  `Add a new task with a title and optional description.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Join all arguments to form the title
		title := strings.Join(args, " ")

		// Add task to database
		task, err := db.AddTask(title, taskDescription)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✓ Task added successfully!\n")
		fmt.Printf("  ID: %d\n", task.ID)
		fmt.Printf("  Title: %s\n", task.Title)
		if task.Description != "" {
			fmt.Printf("  Description: %s\n", task.Description)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Display all tasks with their completion status.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get all tasks from database
		tasks, err := db.GetAllTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error retrieving tasks: %v\n", err)
			os.Exit(1)
		}

		// Check if there are no tasks
		if len(tasks) == 0 {
			fmt.Println("No tasks found. Add one with: taskman add \"Your task\"")
			return
		}

		// Display tasks
		fmt.Println("\nYour Tasks:")
		fmt.Println(strings.Repeat("─", 80))
		for _, task := range tasks {
			// Status indicator
			status := "[ ]"
			if task.Completed {
				status = "[✓]"
			}

			// Format the date
			dateStr := task.CreatedAt.Format("2006-01-02 15:04")

			// Print task info
			fmt.Printf("%s ID: %-3d | %s\n", status, task.ID, task.Title)
			if task.Description != "" {
				fmt.Printf("    Description: %s\n", task.Description)
			}
			fmt.Printf("    Created: %s\n", dateStr)
			fmt.Println(strings.Repeat("─", 80))
		}
		fmt.Printf("\nTotal: %d task(s)\n", len(tasks))
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete [task id]",
	Short: "Mark a task as complete",
	Long:  `Mark a task as completed by providing its ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Complete command - Coming soon!")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "Delete a task",
	Long:  `Remove a task from the database by providing its ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete command - Coming soon!")
	},
}

func init() {
	// Add description flag to add command
	addCmd.Flags().StringVarP(&taskDescription, "description", "d", "", "Task description")

	// Add subcommands to root
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
}

func main() {
	// Initialize database
	if err := db.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Execute CLI commands
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
