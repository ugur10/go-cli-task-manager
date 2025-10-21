# TaskMan - CLI Task Manager

A simple command-line task management application built with Go, demonstrating Go's syntax, idioms, and best practices.

## Features

- ✅ Add tasks with titles and optional descriptions
- ✅ List all tasks with completion status
- ✅ Mark tasks as complete
- ✅ Delete tasks
- ✅ SQLite persistence for data storage
- ✅ Input validation and error handling
- ✅ Clean, formatted CLI output

## Technology Stack

- **Language**: Go
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) - Powerful CLI framework
- **Database**: SQLite (via [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)) - Pure Go SQLite driver
- **Module**: github.com/ugur10/go-cli-task-manager

## Installation

### Prerequisites
- Go 1.21 or later

### Build from Source

```bash
# Clone the repository
git clone https://github.com/ugur10/go-cli-task-manager.git
cd go-cli-task-manager

# Build the application
go build -o taskman

# Run it
./taskman --help
```

## Usage

### Add Tasks

```bash
# Add a simple task
./taskman add "Buy groceries"

# Add a task with description
./taskman add "Learn Go" -d "Study Go programming language fundamentals"
./taskman add "Deploy app" --description "Deploy to production server"
```

### List Tasks

```bash
./taskman list
```

Output example:
```
Your Tasks:
────────────────────────────────────────────────────────────────────────────────
[ ] ID: 3   | Buy groceries
    Created: 2025-10-21 15:30
────────────────────────────────────────────────────────────────────────────────
[✓] ID: 2   | Learn Go
    Description: Study Go programming language fundamentals
    Created: 2025-10-21 14:20
────────────────────────────────────────────────────────────────────────────────

Total: 2 task(s)
```

### Complete Tasks

```bash
# Mark task as complete by ID
./taskman complete 1
```

### Delete Tasks

```bash
# Delete a task by ID
./taskman delete 1
```

### Get Help

```bash
# General help
./taskman --help

# Command-specific help
./taskman add --help
./taskman list --help
./taskman complete --help
./taskman delete --help
```

## Project Structure

```
go-cli-task-manager/
├── main.go           # CLI command definitions and handlers
├── db/
│   ├── db.go         # Database initialization and CRUD operations
│   └── db_test.go    # Unit tests for database functions
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
└── README.md         # This file
```

## Database

Tasks are stored in an SQLite database located at `~/.taskman/tasks.db`. The database is automatically created on first run.

### Schema

```sql
CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL
);
```

## Go Concepts Demonstrated

This project demonstrates various Go programming concepts:

- **Package organization**: Separating concerns into packages (`main`, `db`)
- **Struct definitions**: `Task` struct with typed fields
- **Error handling**: Error wrapping with `fmt.Errorf` and `%w` verb
- **Database operations**: SQL queries, parameterized statements, transactions
- **CLI development**: Using Cobra for command-line interfaces
- **Testing**: Table-driven tests with `testing` package
- **String manipulation**: `strings.Join()`, `strings.Repeat()`
- **Time handling**: `time.Time` and formatting with reference time
- **Type conversion**: `strconv.Atoi()` for string to int
- **Defer pattern**: Resource cleanup with `defer`
- **Slices**: Dynamic arrays and operations
- **Input validation**: Argument validation with Cobra

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests for specific package
go test ./db -v
```

## Development Notes

- Database location: `~/.taskman/tasks.db`
- Binary name: `taskman` (not tracked in git)
- Pure Go implementation (no CGO dependencies)
- Cross-platform compatible

## Future Enhancements

Potential improvements for learning more Go:
- Add task editing functionality
- Implement task filtering (show only completed/incomplete)
- Add task priority levels
- Export/import tasks (JSON/CSV)
- Add due dates and reminders
- Implement task search functionality
- Add color output support

## License

MIT

## Contributing

This is an educational project built to demonstrate Go programming concepts. Feel free to fork and experiment!
