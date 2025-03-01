package tasks

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"skillsrock-test/models"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository { return &Repository{db: db} }

//go:embed sql/select_tasks.sql
var selectTasksSql string

func (r *Repository) SelectTasks() ([]models.Tasks, error) {
	var tasks []models.Tasks

	rows, err := r.db.Query(selectTasksSql)
	if err != nil {
		return []models.Tasks{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Tasks
		if err = rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return []models.Tasks{}, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

//go:embed sql/insert_tasks.sql
var insertTasksSql string

func (r *Repository) InsertTasks(tasks models.Tasks) error {
	_, err := r.db.Exec(insertTasksSql, tasks.Title, tasks.Description, tasks.Status)
	if err != nil {
		return err
	}
	return err
}

//go:embed sql/update_tasks.sql
var updateTasksSql string

func (r *Repository) UpdateTasks(tasks models.Tasks) error {
	_, err := r.db.Exec(updateTasksSql, tasks.Title, tasks.Description, tasks.Status, tasks.Id)
	if err != nil {
		return err
	}
	return err
}

//go:embed sql/delete_tasks.sql
var deleteTasksSql string

func (r *Repository) DeleteTasks(id int) error {
	_, err := r.db.Exec(deleteTasksSql, id)
	if err != nil {
		return err
	}
	return err
}
