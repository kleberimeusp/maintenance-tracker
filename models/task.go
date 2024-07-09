package models

type Task struct {
	ID      int    `db:"id"`
	Summary string `db:"summary"`
	Date    string `db:"date"`
	UserID  int    `db:"user_id"`
}
