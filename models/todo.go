package models

// model or domain
import (
	"context"
)

// Book ...
type User_todo_list struct {
	ID        int64  `json:"id"`
	Task_name string `json:"task_name"`
}

type TodoEntity interface {
	Fetch(ctx context.Context) ([]User_todo_list, error)
	GetByID(ctx context.Context, id int64) (User_todo_list, error)
	Create(ctx context.Context, todo User_todo_list) error
	Update(ctx context.Context, todo User_todo_list, id int64) error
	Delete(ctx context.Context, id int64) error
}

type TodoRepository interface {
	Fetch(ctx context.Context) (res []User_todo_list, err error)
	GetByID(ctx context.Context, id int64) (User_todo_list, error)
	Create(ctx context.Context, todo User_todo_list) error
	Update(ctx context.Context, todo User_todo_list, id int64) error
	Delete(ctx context.Context, id int64) error
}
