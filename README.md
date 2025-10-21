# TaskMan - CLI Task Manager

A simple command-line task management application built with Go, demonstrating Go's syntax, idioms, and best practices.

## Features

- Add tasks with titles and descriptions
- List all tasks with completion status
- Mark tasks as complete
- Delete tasks
- SQLite persistence for data storage

## Technology Stack

- **Language**: Go
- **CLI Framework**: Cobra
- **Database**: SQLite
- **Module**: github.com/ugur10/go-cli-task-manager

## Installation

```bash
go build -o taskman
```

## Usage

```bash
# Show help
./taskman --help

# Add a task
./taskman add "Buy groceries"

# List all tasks
./taskman list

# Complete a task
./taskman complete 1

# Delete a task
./taskman delete 1
```

## Project Status

This project is currently under development. Commands will be implemented incrementally.

## License

MIT
