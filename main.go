package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskman",
	Short: "A simple CLI task manager",
	Long: `TaskMan is a command-line task management application built with Go.
It allows you to add, list, complete, and delete tasks with SQLite persistence.`,
}

var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Long:  `Add a new task with a title and optional description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command - Coming soon!")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `Display all tasks with their completion status.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List command - Coming soon!")
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
	// Add subcommands to root
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(deleteCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
