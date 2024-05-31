package storage

import (
	"database/sql"
	"todo-tui/internal/models"
)

type TasksRepository interface {
	AddTask(task models.Item) error
	RemoveTask(taskId int) error
	GetTasks(status models.Status) ([]models.Item, error)
	GetTask(taskId int) (models.Item, error)
  UpdateTask(task models.Item) error
	CreateNewDb() error
}

type tasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) TasksRepository {
	return &tasksRepository{db}
}

func (r *tasksRepository) CreateNewDb() error {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
      id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
      title       TEXT NOT NULL,
      description TEXT NOT NULL,
      created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      status      INTEGER NOT NULL
    );
  `
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

// status(int): 0 - inprogress | 1 - todo | 2 - done
func (r *tasksRepository) GetTasks(status models.Status) ([]models.Item, error) {
	items := make([]models.Item, 0)
	query := `SELECT id, title, description, status FROM tasks WHERE status = ?`
	rows, err := r.db.Query(query, status)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(
			&item.Id,
			&item.Title,
			&item.Description,
			&item.Status,
		); err != nil {
			return items, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (r *tasksRepository) GetTask(taskId int) (models.Item, error) {
	item := models.Item{}
	query := `SELECT id, title, description, status FROM tasks WHERE is = ?`
	row := r.db.QueryRow(query, taskId)

	if err := row.Scan(&item.Id, &item.Title, &item.Description, &item.Status); err != nil {
		return item, err
	}

	return item, nil
}

func (r *tasksRepository) UpdateTask(task models.Item) error {
	query := `UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?;`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) AddTask(task models.Item) error {
	query := `INSERT INTO tasks (title, description, status) VALUES (?, ?, ?);`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status)

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) RemoveTask(taskId int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := r.db.Exec(query, taskId)

	if err != nil {
		return err
	}

	return nil
}
