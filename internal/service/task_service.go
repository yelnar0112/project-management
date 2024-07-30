package service

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/yelnar0112/project-management/internal/config"
	"github.com/yelnar0112/project-management/internal/domain"
)

func GetAllTasks() (tasks []domain.Task, err error) {
	rows, err := config.DB.Query("SELECT id, title, description, priority, state, assignee, project_id, created_at, completed_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.State, &task.Assignee, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(task *domain.Task) error {
	_, err := config.DB.Exec(
		"INSERT INTO tasks (id, title, description, priority, state, assignee, project_id, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		task.ID, task.Title, task.Description, task.Priority, task.State, task.Assignee, task.ProjectID, task.CreatedAt, task.CompletedAt,
	)
	return err
}

func GetTask(id uuid.UUID) (*domain.Task, error) {
	var task domain.Task
	err := config.DB.QueryRow(
		"SELECT id, title, description, priority, state, assignee, project_id, created_at, completed_at FROM tasks WHERE id = $1", id,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.State, &task.Assignee, &task.ProjectID, &task.CreatedAt, &task.CompletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *domain.Task) error {
	_, err := config.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2, priority = $3, state = $4, assignee = $5, project_id = $6, created_at = $7, completed_at = $8 WHERE id = $9",
		task.Title, task.Description, task.Priority, task.State, task.Assignee, task.ProjectID, task.CreatedAt, task.CompletedAt, task.ID,
	)
	return err
}

func DeleteTask(id uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
