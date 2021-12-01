package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/KennyKur/CRUD_Todo/models"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestTodoRepository_Fetch(t *testing.T) {
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
		mockClosure func(mock sqlmock.Sqlmock)
		fields      fields
		args        args
		wantRes     []models.User_todo_list
		wantErr     bool
	}{
		{
			name: "berhasil mengambil data",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			wantRes: mockTodo,
			wantErr: false,
		},
		{
			name: "gagal mengambil data",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(fmt.Errorf("some error"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			tt.mockClosure(mock)

			m := NewTodoRepository(db)

			gotRes, err := m.Fetch(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoRepository.Fetch() error = %v, wantErr %v", err, tt.wantErr)
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
	mockTodo := []models.User_todo_list{
		{
			ID: 2, Task_name: "Sprint Test",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "task_name"}).
		AddRow(mockTodo[0].ID, mockTodo[0].Task_name)

	query := "SELECT id, task_name FROM user_todo_lists WHERE id = $1"

	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	test_args := args{ctx: context.TODO(), id: 2}
	tests := []struct {
		name        string
		fields      fields
		mockClosure func(mock sqlmock.Sqlmock)
		args        args
		wantRes     models.User_todo_list
		wantErr     bool
	}{
		{
			name: "berhasil mengambil data",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(test_args.id).
					WillReturnRows(rows)
			},
			args:    test_args,
			wantErr: false,
			wantRes: mockTodo[0],
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			tt.mockClosure(mock)

			m := NewTodoRepository(db)

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
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx  context.Context
		todo models.User_todo_list
	}
	data := models.User_todo_list{Task_name: "halo_bandung"}
	data2 := models.User_todo_list{Task_name: "tidur"}
	test_arg := args{ctx: context.TODO(), todo: data}
	test_arg2 := args{ctx: context.TODO(), todo: data2}
	query := "INSERT INTO user_todo_lists(task_name) VALUES ($1)"
	tests := []struct {
		name        string
		fields      fields
		mockClosure func(mock sqlmock.Sqlmock)
		args        args
		wantErr     bool
	}{
		{
			name: "data sukses ditambahkan",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectPrepare(regexp.QuoteMeta(query))
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(test_arg.todo.Task_name).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			args:    test_arg,
			wantErr: false,
		},
		{
			name: "data tidak masuk dalam list todo",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectPrepare(regexp.QuoteMeta(query))
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(test_arg2.todo.Task_name).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			args:    test_arg2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			// Set mock expectations
			tt.mockClosure(mock)

			m := NewTodoRepository(db)
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
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx  context.Context
		todo models.User_todo_list
		id   int64
	}
	data := models.User_todo_list{Task_name: "halo_bandung"}
	test_arg := args{ctx: context.TODO(), todo: data, id: 2}
	query := "UPDATE user_todo_lists SET task_name = $1 WHERE id = $2"
	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		fields      fields
		args        args
		wantErr     bool
	}{
		{
			name: "sukses mengubah data",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectPrepare(regexp.QuoteMeta(query))
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(test_arg.todo.Task_name, test_arg.id).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			args:    test_arg,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			// Set mock expectations
			tt.mockClosure(mock)

			m := NewTodoRepository(db)
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
	type fields struct {
		Conn *sql.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	test_arg := args{ctx: context.TODO(), id: 4}
	query := "DELETE FROM user_todo_lists WHERE id = $1"
	tests := []struct {
		name        string
		mockClosure func(mock sqlmock.Sqlmock)
		fields      fields
		args        args
		wantErr     bool
	}{
		{
			name: "sukses mengubah data",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectPrepare(regexp.QuoteMeta(query))
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(test_arg.id).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			args:    test_arg,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			// Set mock expectations
			tt.mockClosure(mock)

			m := NewTodoRepository(db)
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
