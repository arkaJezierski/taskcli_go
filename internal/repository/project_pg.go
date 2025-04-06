package repository

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/ajezierski/taskcli/internal/db"
	"github.com/ajezierski/taskcli/internal/model"
)

func CreateProject(name string) (string, error) {
	id := uuid.New().String()
	_, err := db.DB.Exec("INSERT INTO projects (id, name) VALUES ($1, $2)", id, name)
	return id, err
}

func GetProject(id string) (*model.Project, error) {
	row := db.DB.QueryRow("SELECT id, name FROM projects WHERE id = $1", id)
	var p model.Project
	if err := row.Scan(&p.ID, &p.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func UpdateProject(id string, name string) error {
	_, err := db.DB.Exec("UPDATE projects SET name = $1 WHERE id = $2", name, id)
	return err
}

func DeleteProject(id string) error {
	_, err := db.DB.Exec("DELETE FROM projects WHERE id = $1", id)
	return err
}

func ListProjects() ([]model.Project, error) {
	rows, err := db.DB.Query("SELECT id, name FROM projects ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}
