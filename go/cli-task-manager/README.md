# CLI Task Manager

A simple command-line tool built with Go and Cobra to manage your daily tasks. It uses a local SQLite database (`tasks.db`) to store and manage tasks.

## Features

* **Add new tasks**: Add tasks with a description and an optional priority level (`high`, `medium`, `low`).
* **List tasks**: View all your pending tasks. You can also view all tasks, including completed ones.
* **Complete tasks**: Mark tasks as completed by their ID.
* **Delete tasks**: Remove tasks from your list by their ID.

## Installation

1. **Prerequisites**: Ensure you have Go (version 1.24 or later) installed on your system.

2. **Build from source**:

    ```bash
    # Clone the repository (or navigate to the project directory)
    # go get -d ./... # To install dependencies
    go build -o task
    ```

    This will create an executable file named `task` in the project directory.

## Usage

All commands are run from the root of the application.

### Add a Task

Create a new task with a description. The default priority is `medium`.

```bash
./task add "My first task"
```

You can specify a priority using the `-p` or `--priority` flag.

```bash
./task add "An important task" --priority high
```

### List Tasks

To see all your pending tasks:

```bash
./task list
```

To see all tasks, including those that are already completed, use the `-a` or `--all` flag:

```bash
./task list --all
```

### Complete a Task

Mark a task as complete using its ID.

```bash
./task complete 1
```

### Delete a Task

Delete a task from the list using its ID.

```bash
./task delete 1
```

## Project Structure

```
cli-task-manager/
├── cmd/          # Cobra command definitions (add, list, etc.)
├── pkg/
│   ├── db/       # Database logic (SQLite)
│   ├── models/   # Data models (Task)
│   └── utils/    # Utility functions
├── go.mod
├── go.sum
└── main.go       # Main application entry point
```


## Dependencies

* [github.com/spf13/cobra](https://github.com/spf13/cobra): A popular library for creating powerful modern CLI applications.
* [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3): SQLite3 driver for Go.
