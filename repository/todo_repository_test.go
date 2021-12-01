package repository

import (
	"context"
	"database/sql"
	"reflect"
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
			name: "berhasil mengambil data",
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
