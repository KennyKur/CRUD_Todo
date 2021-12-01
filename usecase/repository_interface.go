package usecase

import (
	"context"

	"github.com/KennyKur/CRUD_Todo/models"
)

type TodoRepositoryInterface interface {
	Fetch(ctx context.Context) (res []models.User_todo_list, err error)
	GetByID(ctx context.Context, id int64) (models.User_todo_list, error)
	Create(ctx context.Context, todo models.User_todo_list) error
	Update(ctx context.Context, todo models.User_todo_list, id int64) error
	Delete(ctx context.Context, id int64) error
}
