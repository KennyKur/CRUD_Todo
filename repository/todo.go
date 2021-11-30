package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/KennyKur/CRUD_Todo/models"
)

var list_not_todo = []string{
	"cuti",
	"berenang",
	"tidur",
}

type TodoRepository struct {
	Conn *sql.DB
}

func NewTodoRepository(Conn *sql.DB) models.TodoRepository {
	return &TodoRepository{Conn}
}

func (m *TodoRepository) Fetch(ctx context.Context) (res []models.User_todo_list, err error) {
	rows, err := m.Conn.Query("SELECT id, task_name FROM user_todo_lists")
	if err != nil {
		return
	}
	var todos []models.User_todo_list
	for rows.Next() {
		var todo models.User_todo_list
		rows.Scan(&todo.ID, &todo.Task_name)
		todos = append(todos, todo)
	}

	defer rows.Close()
	return todos, nil
}

func (m *TodoRepository) GetByID(ctx context.Context, id int64) (res models.User_todo_list, err error) {
	var todo models.User_todo_list
	row := m.Conn.QueryRow("SELECT id, task_name FROM user_todo_lists WHERE id = $1", id)
	err = row.Scan(&todo.ID, &todo.Task_name)
	if err != nil {
		return
	}
	return todo, nil
}

func (m TodoRepository) Create(ctx context.Context, todo models.User_todo_list) error {
	tx, err := m.Conn.Begin()
	for _, b := range list_not_todo {
		if b == todo.Task_name {
			err := errors.New("task tidak bisa dimasukan")
			return err
		}
	}
	if err != nil {
		return err
	}
	{
		stmt, err := tx.Prepare("INSERT INTO user_todo_lists(task_name) VALUES ($1)")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(todo.Task_name); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
func (m *TodoRepository) Update(ctx context.Context, todo models.User_todo_list, id int64) error {
	tx, err := m.Conn.Begin()
	for _, b := range list_not_todo {
		if b == todo.Task_name {
			err := errors.New("task tidak bisa dimasukan")
			return err
		}
	}
	if err != nil {
		return err
	}
	{
		stmt, err := tx.Prepare("UPDATE user_todo_lists SET task_name = $1 WHERE id = $2")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(todo.Task_name, id); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (m *TodoRepository) Delete(ctx context.Context, id int64) error {
	tx, err := m.Conn.Begin()
	if err != nil {
		return err
	}
	{
		stmt, err := tx.Prepare("DELETE FROM user_todo_lists WHERE id = $1")
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()
		if _, err := stmt.Exec(id); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil

}
