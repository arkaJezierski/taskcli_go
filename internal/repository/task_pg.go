package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/ajezierski/taskcli/internal/db"
	"github.com/ajezierski/taskcli/internal/model"
)

func CreateTaskToProject(projectIdentifier, title string) (string, error) {
	id := uuid.New().String()
	var projectID string

	if _, err := uuid.Parse(projectIdentifier); err == nil {
		projectID = projectIdentifier
	} else {
		err := db.DB.QueryRow("SELECT id FROM projects WHERE name = $1", projectIdentifier).Scan(&projectID)
		if err != nil {
			return "", fmt.Errorf("project not found: %v", err)
		}
	}

	_, err := db.DB.Exec("INSERT INTO tasks (id, project_id, title) VALUES ($1, $2, $3)", id, projectID, title)
	if err != nil {
		return "", fmt.Errorf("insert failed: %v", err)
	}

	return id, nil
}

func ListTasks(projectID string) ([]model.Task, error) {
	var rows *sql.Rows
	var err error

	if projectID != "" {
		if _, err := uuid.Parse(projectID); err == nil {
			rows, err = db.DB.Query(`
				SELECT t.id, p.name as project_name, t.title, t.done
				FROM tasks t
				JOIN projects p ON t.project_id = p.id
				WHERE t.project_id = $1
			`, projectID)
		} else {
			rows, err = db.DB.Query(`
				SELECT t.id, p.name as project_name, t.title, t.done
				FROM tasks t
				JOIN projects p ON t.project_id = p.id
				WHERE p.name = $1
			`, projectID)
		}
	} else {
		rows, err = db.DB.Query("SELECT t.id, p.name as project_name, t.title, t.done FROM tasks t join projects p on t.project_id = p.id")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func MarkTaskDone(id string) error {
	_, err := db.DB.Exec("UPDATE tasks SET done = TRUE WHERE id = $1", id)
	return err
}

func DeleteTask(id string) error {
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
