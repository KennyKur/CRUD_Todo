package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/KennyKur/CRUD_Todo/models"
	gomock "github.com/golang/mock/gomock"
)

func TestTodoUsecase_Fetch(t *testing.T) {
	mockTodos := []models.User_todo_list{
		{
			ID: 1, Task_name: "Belajar",
		},
		{
			ID: 2, Task_name: "Sprint Test",
		},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := NewMockTodoRepositoryInterface(ctrl)
	type fields struct {
		todoRepo TodoRepositoryInterface
	}
	type args struct {
		c context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFN  func(args)
		wantRes []models.User_todo_list
		wantErr bool
	}{
		{
			name: "success to  get data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c: context.Background(),
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					Fetch(a.c).
					Return(mockTodos, nil)
			},
			wantRes: mockTodos,
			wantErr: false,
		},
		{
			name: "failed to get data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c: context.Background(),
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					Fetch(a.c).
					Return(nil, errors.New("gagal mengambil data"))
			},
			wantRes: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			tt.mockFN(tt.args)
			a := &TodoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			gotRes, err := a.Fetch(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("TodoUsecase.Fetch() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestTodoUsecase_GetByID(t *testing.T) {
	mockTodo := models.User_todo_list{ID: 4, Task_name: "daily"}
	mockTodoErr := models.User_todo_list{ID: 0, Task_name: ""}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := NewMockTodoRepositoryInterface(ctrl)
	type fields struct {
		todoRepo TodoRepositoryInterface
	}
	type args struct {
		c  context.Context
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFN  func(args)
		wantRes models.User_todo_list
		wantErr bool
	}{
		{
			name: "success to get data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:  context.Background(),
				id: 4,
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					GetByID(a.c, a.id).
					Return(mockTodo, nil)
			},
			wantRes: mockTodo,
			wantErr: false,
		},
		{
			name: "failed to get data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:  context.Background(),
				id: 4,
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					GetByID(a.c, a.id).
					Return(mockTodoErr, errors.New("gagal mendapatkan data"))
			},
			wantRes: mockTodoErr,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFN(tt.args)
			a := &TodoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			gotRes, err := a.GetByID(tt.args.c, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("TodoUsecase.GetByID() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestTodoUsecase_Create(t *testing.T) {
	mockTodo := models.User_todo_list{ID: 1, Task_name: "daily"}
	mockTodo2 := models.User_todo_list{ID: 1, Task_name: "tidur"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := NewMockTodoRepositoryInterface(ctrl)
	type fields struct {
		todoRepo TodoRepositoryInterface
	}
	type args struct {
		c    context.Context
		todo models.User_todo_list
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFN  func(args)
		wantErr bool
	}{
		{
			name: "success to add data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:    context.Background(),
				todo: mockTodo,
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					Create(a.c, a.todo).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to add data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:    context.Background(),
				todo: mockTodo2,
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					Create(a.c, a.todo).
					Return(errors.New("Task tidak valid"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFN(tt.args)
			a := &TodoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			if err := a.Create(tt.args.c, tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoUsecase_Update(t *testing.T) {
	mockTodo := models.User_todo_list{Task_name: "mengerjakan nxt"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := NewMockTodoRepositoryInterface(ctrl)
	type fields struct {
		todoRepo TodoRepositoryInterface
	}
	type args struct {
		c    context.Context
		todo models.User_todo_list
		id   int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFN  func(args)
		wantErr bool
	}{
		{
			name: "success to update data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:    context.Background(),
				todo: mockTodo,
				id:   4,
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					Update(a.c, a.todo, a.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to update data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:    context.Background(),
				todo: mockTodo,
				id:   10,
			},
			mockFN: func(a args) {
				mockUC.EXPECT().
					Update(a.c, a.todo, a.id).
					Return(errors.New("data not found"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFN(tt.args)
			a := &TodoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			if err := a.Update(tt.args.c, tt.args.todo, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUC := NewMockTodoRepositoryInterface(ctrl)
	type fields struct {
		todoRepo TodoRepositoryInterface
	}
	type args struct {
		c  context.Context
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFn  func(args)
		wantErr bool
	}{
		{
			name: "success to delete data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:  context.Background(),
				id: 4,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().
					Delete(a.c, a.id).
					Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to delete data",
			fields: fields{
				todoRepo: mockUC,
			},
			args: args{
				c:  context.Background(),
				id: 5,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().
					Delete(a.c, a.id).
					Return(errors.New("Unexpected error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			a := &TodoUsecase{
				todoRepo: tt.fields.todoRepo,
			}
			if err := a.Delete(tt.args.c, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
