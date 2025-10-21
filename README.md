# TaskMan - CLI Task Manager 📝

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen.svg)](https://github.com/ugur10/go-cli-task-manager)

A simple, elegant command-line task management application built with Go. Perfect for learning Go's syntax, idioms, and best practices through a real-world project.

## ✨ Quick Start

```bash
# Clone and build
git clone https://github.com/ugur10/go-cli-task-manager.git
cd go-cli-task-manager
go build -o taskman

# Start using it!
./taskman add "Learn Go" -d "Build awesome CLI tools"
./taskman list
```

## 🚀 Features

- ✅ **Create** tasks with titles and optional descriptions
- ✅ **Read** and list all tasks with completion status
- ✅ **Update** tasks by marking them complete
- ✅ **Delete** tasks you no longer need
- 💾 **Persistent storage** with SQLite
- 🛡️ **Input validation** and comprehensive error handling
- 🎨 **Clean, formatted** CLI output
- ✔️ **Zero dependencies** (pure Go, no CGO required)
- 🧪 **Fully tested** with comprehensive test suite

## 🛠️ Technology Stack

| Component | Technology |
|-----------|------------|
| **Language** | [Go 1.21+](https://go.dev/) |
| **CLI Framework** | [Cobra](https://github.com/spf13/cobra) |
| **Database** | SQLite ([modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)) |
| **Testing** | Go's built-in `testing` package |

## 📦 Installation

### Prerequisites
- [Go 1.21 or later](https://go.dev/dl/)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/ugur10/go-cli-task-manager.git
cd go-cli-task-manager

# Build the application
go build -o taskman

# Optionally, move to your PATH
sudo mv taskman /usr/local/bin/  # macOS/Linux
```

## 📖 Usage

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

## 📁 Project Structure

```
go-cli-task-manager/
├── main.go           # CLI command definitions and handlers
├── db/
│   ├── db.go         # Database initialization and CRUD operations
│   └── db_test.go    # Unit tests for database functions
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
├── LICENSE           # MIT License
├── .gitignore        # Git ignore rules
└── README.md         # This file
```

## 🗄️ Database

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

## 🎓 Go Concepts Demonstrated

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

## 🧪 Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Run tests for specific package
go test ./db -v
```

## 💡 Development Notes

- 📍 **Database location**: `~/.taskman/tasks.db`
- 🔨 **Binary name**: `taskman` (not tracked in git)
- 🚫 **No CGO**: Pure Go implementation for easy cross-compilation
- 🌍 **Cross-platform**: Works on macOS, Linux, and Windows
- 🧹 **Clean code**: Well-commented and follows Go best practices

## 🔮 Future Enhancements

Potential improvements for learning more Go:

- [ ] Task editing functionality
- [ ] Task filtering (show only completed/incomplete)
- [ ] Task priority levels
- [ ] Export/import tasks (JSON/CSV)
- [ ] Due dates and reminders
- [ ] Task search functionality
- [ ] Color output support with [fatih/color](https://github.com/fatih/color)
- [ ] Configuration file support
- [ ] Task tags/categories

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

This is an educational project built to demonstrate Go programming concepts. Contributions, issues, and feature requests are welcome!

Feel free to:
- 🐛 Report bugs
- 💡 Suggest new features
- 🔧 Submit pull requests
- ⭐ Star the repo if you find it useful!

## 👨‍💻 Author

**Ugur**

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Uses [modernc.org/sqlite](https://modernc.org/sqlite) for pure Go SQLite
- Inspired by the need for simple, effective task management tools
