package repository

import (
	"fmt"
	"github.com/TrackTasks/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) CreateTask(newTask models.Task) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createTaskQuery := fmt.Sprintf("INSERT INTO tasks (description,status) VALUES ($1,$2) RETURNING id")
	row := tx.QueryRow(createTaskQuery, newTask.Description, newTask.Status)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
func (r *TaskPostgres) UpdateTask(id int, task models.Task) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if task.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, task.Description)
		argId++
	}
	if task.Status != "" {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, task.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ",")
	//query := fmt.Sprintf("UPDATE tasks SET %s WHERE id = $1", setQuery, args)
	query := "UPDATE tasks SET "
	query = query + setQuery
	query += fmt.Sprintf(" WHERE id = $%d", argId)

	args = append(args, id)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
func (r *TaskPostgres) DeleteTask(id int, status string) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if id != 0 {
		setValues = append(setValues, fmt.Sprintf("id=$%d", argId))
		args = append(args, id)
		argId++
	}
	if status != "" {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, status)
		argId++
	}
	setQuery := strings.Join(setValues, "AND")
	query := "DELETE FROM tasks WHERE" + setQuery

	res, err := r.db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete people: %s", err)
	}
	return res.RowsAffected()
}
func (r *TaskPostgres) GetTask(taskIdInt int, status string) ([]models.Task, error) {
	var tasks []models.Task
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if taskIdInt != 0 {
		setValues = append(setValues, fmt.Sprintf("id=$%d", argId))
		args = append(args, taskIdInt)
		argId++
	}

	if status != "" {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, status)
		argId++
	}
	query := fmt.Sprintf("SELECT id, description, status FROM tasks WHERE 1 = 1 ")
	if len(setValues) > 0 {
		query += " AND " + strings.Join(setValues, " AND ")
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query people: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}
