// models/model.go
package models

import (
	"database/sql"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// CreateTask saves a new task to the database
func CreateTask(db *sql.DB, task *Task) error {
	_, err := db.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", task.Title, task.Completed)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTask removes a task from the database
func DeleteTask(db *sql.DB, taskID string) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id=?", taskID)
	if err != nil {
		return err
	}
	return nil
}
