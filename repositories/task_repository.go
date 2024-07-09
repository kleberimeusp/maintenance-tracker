package repositories

import (
	"maintenance-tracker/models"

	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) CreateTask(task models.Task) error {
	_, err := r.db.NamedExec(`INSERT INTO tasks (summary, date, user_id) VALUES (:summary, :date, :user_id)`, &task)
	return err
}

func (r *TaskRepository) GetTasksByUser(userID int) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Select(&tasks, `SELECT * FROM tasks WHERE user_id = ?`, userID)
	return tasks, err
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Select(&tasks, `SELECT * FROM tasks`)
	return tasks, err
}

func (r *TaskRepository) DeleteTask(taskID int) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = ?`, taskID)
	return err
}
