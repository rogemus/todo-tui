package tui

import "database/sql"

type TasksRepository interface {
	CreateNewDb() error
	AddTask(task task) error
	RemoveTask(taskId string) error
	UpdateTask(task task) error
	GetTask(taskId int) (task, error)
	GetTasks() ([]task, error)
	GetTasksByStatus(status bool) ([]task, error)
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
      done        BOOLEAN NOT NULL
    );
  `
	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (r *tasksRepository) GetTasks() ([]task, error) {
	items := make([]task, 0)
	query := `SELECT id, title, description, done FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return items, err
	}

	defer rows.Close()

	for rows.Next() {
		var id, title, desc string
		var done bool

		if err := rows.Scan(&id, &title, &desc, &done); err != nil {
			return items, err
		}
		items = append(items, NewTask(id, title, desc, done))
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (r *tasksRepository) GetTasksByStatus(status bool) ([]task, error) {
	items := make([]task, 0)
	query := `SELECT id, title, description, done FROM tasks WHERE done = ?`

	rows, err := r.db.Query(query, status)
	if err != nil {
		return items, err
	}

	defer rows.Close()

	for rows.Next() {
		var id, title, desc string
		var done bool

		if err := rows.Scan(&id, &title, &desc, &done); err != nil {
			return items, err
		}
		items = append(items, NewTask(id, title, desc, done))
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (r *tasksRepository) GetTask(taskId int) (task, error) {
	item := task{}
	query := `SELECT id, title, description, done FROM tasks WHERE is = ?`
	row := r.db.QueryRow(query, taskId)
	var id, title, desc string
	var done bool

	if err := row.Scan(&id, &title, &desc, &done); err != nil {
		return item, err
	}

	item = NewTask(id, title, desc, done)
	return item, nil
}

func (r *tasksRepository) UpdateTask(task task) error {
	query := `UPDATE tasks SET title = ?, description = ?, done = ? WHERE id = ?`
	_, err := r.db.Exec(query, task.Title(), task.Description(), task.Done(), task.Id())

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) AddTask(task task) error {
	query := `INSERT INTO tasks (title, description, done) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, task.Title(), task.Description(), false)

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) RemoveTask(taskId string) error {
	query := `DELETE FROM tasks WHERE id = ?`
	_, err := r.db.Exec(query, taskId)

	if err != nil {
		return err
	}

	return nil
}
