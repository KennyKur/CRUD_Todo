package entity

import (
	"context"

	"github.com/KennyKur/CRUD_Todo/models"
)

type TodoEntity struct {
	todoRepo models.TodoRepository
}

func NewTodoEntity(a models.TodoRepository) models.TodoEntity {
	return &TodoEntity{
		todoRepo: a,
	}
}

func (a *TodoEntity) Fetch(c context.Context) (res []models.User_todo_list, err error) {
	res, err = a.todoRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return

}

func (a *TodoEntity) GetByID(c context.Context, id int64) (res models.User_todo_list, err error) {
	res, err = a.todoRepo.GetByID(c, id)
	return
}

func (a *TodoEntity) Create(c context.Context, todo models.User_todo_list) error {
	err := a.todoRepo.Create(c, todo)
	if err != nil {
		return err
	}

	return nil
}

func (a *TodoEntity) Update(c context.Context, todo models.User_todo_list, id int64) error {
	err := a.todoRepo.Update(c, todo, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *TodoEntity) Delete(c context.Context, id int64) error {
	err := a.todoRepo.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
