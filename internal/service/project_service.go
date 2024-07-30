package service

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/yelnar0112/project-management/internal/config"
	"github.com/yelnar0112/project-management/internal/domain"
)

func GetAllProjects() (projects []domain.Entity, err error) {
	rows, err := config.DB.Query("SELECT id, title, description, start_date, end_date, manager_id FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project domain.Entity
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func CreateProject(project *domain.Entity) error {
	_, err := config.DB.Exec(
		"INSERT INTO projects (id, title, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5, $6)",
		project.ID, project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID,
	)
	return err
}

func GetProject(id uuid.UUID) (*domain.Entity, error) {
	var project domain.Entity
	err := config.DB.QueryRow(
		"SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE id = $1", id,
	).Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No rows found, return nil instead of an error
		}
		return nil, err
	}
	return &project, nil
}

func UpdateProject(project *domain.Entity) error {
	_, err := config.DB.Exec(
		"UPDATE projects SET title = $1, description = $2, start_date = $3, end_date = $4, manager_id = $5 WHERE id = $6",
		project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID, project.ID,
	)
	return err
}

func DeleteProject(id uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM projects WHERE id = $1", id)
	return err
}
