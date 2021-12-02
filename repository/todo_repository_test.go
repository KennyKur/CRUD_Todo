package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/KennyKur/CRUD_Todo/models"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestTodoRepository_Fetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockTodo := []models.User_todo_list{
		{
			ID: 1, Task_name: "Belajar",
		},
		{
			ID: 2, Task_name: "Sprint Test",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "task_name"}).
		AddRow(mockTodo[0].ID, mockTodo[0].Task_name).
		AddRow(mockTodo[1].ID, mockTodo[1].Task_name)

	query := "SELECT id, task_name FROM user_todo_lists"
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		mockClosure func(mock sqlmock.Sqlmock)
		wantRes     []models.User_todo_list
		wantErr     bool
	}{
		{
			name: "success to get data",
			fields: fields{
				Conn: db,
			},
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			wantRes: mockTodo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.mockClosure(mock)
			m := &TodoRepository{
				Conn: tt.fields.Conn,
			}
			gotRes, err := m.Fetch(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			mock.ExpectClose()
			// Explicit closing instead of deferred in order to check ExpectationsWereMet
			if err = db.Close(); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("TodoRepository.Fetch() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestTodoRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	var id int64 = 5
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockTodo := models.User_todo_list{
		ID: 5, Task_name: "Belajar",
	}
	rows := sqlmock.NewRows([]string{"id", "task_name"}).
		AddRow(mockTodo.ID, mockTodo.Task_name)

	query := "SELECT id, task_name FROM user_todo_lists WHERE id = $1"
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		mockClosure func(mock sqlmock.Sqlmock, a args)
		wantRes     models.User_todo_list
		wantErr     bool
	}{
		{
			name: "success to get data",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx: context.Background(),
				id:  id,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(a.id).
					WillReturnRows(rows)
			},
			wantRes: mockTodo,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			tt.mockClosure(mock, tt.args)
			m := &TodoRepository{
				Conn: tt.fields.Conn,
			}
			gotRes, err := m.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			mock.ExpectClose()
			// Explicit closing instead of deferred in order to check ExpectationsWereMet
			if err = db.Close(); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("TodoRepository.GetByID() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestTodoRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	data := models.User_todo_list{Task_name: "daily_harian"}
	data2 := models.User_todo_list{Task_name: "tidur"}
	query := "INSERT"
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx  context.Context
		todo models.User_todo_list
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		mockClosure func(mock sqlmock.Sqlmock, a args)
		wantErr     bool
	}{
		{
			name: "success to add data",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx:  context.Background(),
				todo: data,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.todo.Task_name).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "failed to create data (invalid data)",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx:  context.Background(),
				todo: data2,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.todo.Task_name).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: true,
		},
		{
			name: "failed to create data (query error)",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx:  context.Background(),
				todo: data,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.todo.Task_name).
					WillReturnError(fmt.Errorf("some error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := TodoRepository{
				Conn: tt.fields.Conn,
			}
			tt.mockClosure(mock, tt.args)
			if err := m.Create(tt.args.ctx, tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			mock.ExpectClose()
			// Explicit closing instead of deferred in order to check ExpectationsWereMet
			if err = db.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestTodoRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	data := models.User_todo_list{Task_name: "halo_bandung"}
	data2 := models.User_todo_list{Task_name: "tidur"}
	var id int64 = 2
	query := "UPDATE"
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx  context.Context
		todo models.User_todo_list
		id   int64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		mockClosure func(mock sqlmock.Sqlmock, a args)
		wantErr     bool
	}{
		{
			name: "success update data",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx:  context.Background(),
				todo: data,
				id:   id,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.todo.Task_name, a.id).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "failed update data (invalid data)",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx:  context.Background(),
				todo: data2,
				id:   id,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.todo.Task_name, a.id).WillReturnError(fmt.Errorf("some error"))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "failed update data (sql error)",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx:  context.Background(),
				todo: data,
				id:   id,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.todo.Task_name, a.id).WillReturnError(fmt.Errorf("some error"))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TodoRepository{
				Conn: tt.fields.Conn,
			}
			tt.mockClosure(mock, tt.args)
			if err := m.Update(tt.args.ctx, tt.args.todo, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			mock.ExpectClose()
			// Explicit closing instead of deferred in order to check ExpectationsWereMet
			if err = db.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestTodoRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	var id int64 = 2
	query := "DELETE"
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		mockClosure func(mock sqlmock.Sqlmock, a args)
		wantErr     bool
	}{
		{
			name: "success delete data",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx: context.Background(),
				id:  id,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.id).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "failed to delete data",
			fields: fields{
				Conn: db,
			},
			args: args{
				ctx: context.Background(),
				id:  id,
			},
			mockClosure: func(mock sqlmock.Sqlmock, a args) {
				mock.ExpectBegin()
				mock.ExpectPrepare(query)
				mock.ExpectExec(query).
					WithArgs(a.id).WillReturnError(fmt.Errorf("some error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TodoRepository{
				Conn: tt.fields.Conn,
			}
			tt.mockClosure(mock, tt.args)
			if err := m.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			mock.ExpectClose()
			// Explicit closing instead of deferred in order to check ExpectationsWereMet
			if err = db.Close(); err != nil {
				t.Error(err)
			}
		})
	}
}
