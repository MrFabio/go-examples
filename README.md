# Go Learning Examples

A Terminal User Interface (TUI) application that provides interactive examples and demonstrations of various Go programming concepts. This project serves as both a learning resource and a reference implementation of Go language features. Based on the [excellent tutorial by Derek Banas](https://www.youtube.com/watch?v=YzLrWHZa-Kc).

This repository is part of my Go learning journey, where I'm documenting and implementing various Go concepts through practical examples. Each example is designed to be both educational and immediately runnable, making it easier to understand Go's features in action.

## 🚀 Features

- Interactive TUI interface for exploring Go examples
- Covers some common Go language features
- Displays code with syntax highlighting and its output
- Organized examples by category

## 📚 Example Categories

- **Data Types**
  - Runes
  - Strings
  - Integers
  - Arrays
  - Slices
  - Maps
  - Generics
  - Structs
  - Composition

- **Functions & Control Flow**
  - Time operations
  - Loops
  - Function calls
  - I/O operations
  - CLI interactions

- **Language Features**
  - Accessors
  - Interfaces
  - Closures
  - Recursion
  - Regular Expressions

- **Concurrency**
  - Goroutines
  - Simple Channels
  - Channels with Mutex
  - Concurrent programming patterns

## 🌐 Web Application

The `webapp` package contains a simple Todo application that demonstrates:

- Basic web server setup using Go's standard library
- RESTful API implementation
- CRUD operations for todo items
- Simple file storage
- Basic HTML templates
- HTTP routing and handlers

To run the web application:

```bash
cd webapp
go run main.go
```

The server will start on `http://localhost:8080`

## 📦 Installation

1. Install dependencies:

```bash
go mod download
```

2. Build the project:

```bash
go build
```

This will create an executable named `hello` in the current directory.

## 🎮 Usage

Run the application:

```bash
./hello
```

Use the following controls:

- ↑/↓: Navigate through examples
- Enter: Select and run an example
- q: Quit the application
- h: Show help

## 🧩 Dependencies used

- [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles) - Provides the core UI components like lists and viewports for building the interactive interface from [`charmbracelet/bubbletea`](https://github.com/charmbracelet/bubbletea)
- [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss) - Used for styling the terminal UI with colors, borders, and layout management
- [alecthomas/chroma](https://github.com/alecthomas/chroma) - Provides syntax highlighting for code examples in the terminal

## 📁 Project Structure

``` c#
.
├── main.go          # Application entry point
├── tui.go           # Terminal UI implementation
├── datatypes/       # Data type examples
├── functions/       # Function examples
├── lang/            # Language feature examples
│   └── concurrency/ # Concurrency examples
├── utils/           # Utility functions
└── webapp/          # Web application that has a simple Todo app
```
