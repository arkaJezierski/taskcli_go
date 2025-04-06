package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ajezierski/taskcli/internal/db"
	"github.com/ajezierski/taskcli/internal/model"
)

func CreateTaskToProject(projectID, title string) (string, error) {
	id := uuid.New().String()
	_, err := db.DB.Exec("INSERT INTO tasks (id, project_id, title) VALUES ($1, $2, $3)", id, projectID, title)
	return id, err
}

func ListTasks(projectID string) ([]model.Task, error) {
	var rows *sql.Rows
	var err error

	if projectID != "" {
		rows, err = db.DB.Query("SELECT t.id, p.name as project_name, t.title, t.done FROM tasks t join projects p on t.project_id = p.id WHERE project_id = $1", projectID)
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
