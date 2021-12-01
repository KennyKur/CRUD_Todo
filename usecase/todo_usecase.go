package usecase

import (
	"context"

	"github.com/KennyKur/CRUD_Todo/handler"
	"github.com/KennyKur/CRUD_Todo/models"
)

type TodoUsecase struct {
	todoRepo TodoRepositoryInterface
}

func NewTodoUsecase(a TodoRepositoryInterface) handler.TodoUsecaseInterface {
	return &TodoUsecase{
		todoRepo: a,
	}
}

func (a *TodoUsecase) Fetch(c context.Context) (res []models.User_todo_list, err error) {
	res, err = a.todoRepo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return

}

func (a *TodoUsecase) GetByID(c context.Context, id int64) (res models.User_todo_list, err error) {
	res, err = a.todoRepo.GetByID(c, id)
	return
}

func (a *TodoUsecase) Create(c context.Context, todo models.User_todo_list) error {
	err := a.todoRepo.Create(c, todo)
	if err != nil {
		return err
	}

	return nil
}

func (a *TodoUsecase) Update(c context.Context, todo models.User_todo_list, id int64) error {
	err := a.todoRepo.Update(c, todo, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *TodoUsecase) Delete(c context.Context, id int64) error {
	err := a.todoRepo.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
