package db

import (
	"database/sql"
	"time"

	"task/pkg/models"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	conn *sql.DB
}

func NewSQLite(dbFile string) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Sqlite{
		conn: db,
	}, nil
}

func (db *Sqlite) CreateTasksTable() error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"description" TEXT,
		"priority" TEXT,
		"completed" TEXT,
		"created_at" DATETIME,
		"completed_at" DATETIME
	);`

	statement, err := db.conn.Prepare(createTableSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec()
	return err
}

func (db *Sqlite) InsertTask(task *models.Task) (int64, error) {
	insertSQL := "INSERT INTO tasks(description, priority, completed, created_at, completed_at) VALUES (?, ?, ?, ?, ?)"
	statement, err := db.conn.Prepare(insertSQL)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(task.Description, task.Priority, task.State, task.CreatedAt, task.CompletedAt)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *Sqlite) GetAllTasks() ([]models.Task, error) {
	rows, err := db.conn.Query("SELECT id, description, priority, completed, created_at, completed_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		var completedAt sql.NullTime
		if err := rows.Scan(&task.ID, &task.Description, &task.Priority, &task.State, &task.CreatedAt, &completedAt); err != nil {
			return nil, err
		}
		if completedAt.Valid {
			task.CompletedAt = &completedAt.Time
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (db *Sqlite) CompleteTask(id int) error {
	updateSQL := "UPDATE tasks SET completed = ?, completed_at = ? WHERE id = ?"
	statement, err := db.conn.Prepare(updateSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(models.StateCompleted, time.Now(), id)
	return err
}

func (db *Sqlite) DeleteTask(id int) error {
	deleteSQL := "DELETE FROM tasks WHERE id = ?"
	statement, err := db.conn.Prepare(deleteSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
