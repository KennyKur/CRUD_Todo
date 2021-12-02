package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KennyKur/CRUD_Todo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestTodoHandler_FindTodos(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/Todo", nil)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUC := NewMockTodoUsecaseInterface(ctrl)
	type fields struct {
		TodoUsecase TodoUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockFn func(args)
	}{
		{
			name: "berhasil mengambil data",
			fields: fields{
				TodoUsecase: mockUC,
			},
			args: args{
				c: ctx,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().
					Fetch(a.c.Request.Context())

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			a := &TodoHandler{
				TodoUsecase: tt.fields.TodoUsecase,
			}
			a.FindTodos(tt.args.c)
		})
	}
}

func TestTodoHandler_FindTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/Todo/:id", nil)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUC := NewMockTodoUsecaseInterface(ctrl)
	type fields struct {
		TodoUsecase TodoUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockFn func(args)
	}{
		{
			name: "sukses mengambil data",
			fields: fields{
				TodoUsecase: mockUC,
			},
			args: args{
				c: ctx,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().
					GetByID(a.c.Request.Context(), gomock.Any())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			a := &TodoHandler{
				TodoUsecase: tt.fields.TodoUsecase,
			}
			a.FindTodo(tt.args.c)
		})
	}
}

func TestTodoHandler_CreateTodo(t *testing.T) {

	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/Todos", nil)
	mockTodo := models.User_todo_list{Task_name: "task"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUC := NewMockTodoUsecaseInterface(ctrl)
	type fields struct {
		TodoUsecase TodoUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockFn func(args)
	}{
		{
			name: "sukses menambahkan data",
			fields: fields{
				TodoUsecase: mockUC,
			},
			args: args{
				c: ctx,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().
					Create(a.c.Request.Context(), mockTodo).Return(nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			a := &TodoHandler{
				TodoUsecase: tt.fields.TodoUsecase,
			}
			a.CreateTodo(tt.args.c)
		})
	}
}

func TestTodoHandler_UpdateTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodPatch, "/Todo/update/:id", nil)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUC := NewMockTodoUsecaseInterface(ctrl)

	type fields struct {
		TodoUsecase TodoUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockFn func(args)
	}{
		{
			name: "sukses mengubah data",
			fields: fields{
				TodoUsecase: mockUC,
			},
			args: args{
				c: ctx,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().
					Update(a.c.Request.Context(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			a := &TodoHandler{
				TodoUsecase: tt.fields.TodoUsecase,
			}
			a.UpdateTodo(tt.args.c)
		})
	}
}

func TestTodoHandler_DeleteTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodDelete, "/Todo/delete/:id", nil)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUC := NewMockTodoUsecaseInterface(ctrl)
	type fields struct {
		TodoUsecase TodoUsecaseInterface
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mockFn func(args)
	}{
		{
			name: "sukses menghapus data",
			fields: fields{
				TodoUsecase: mockUC,
			},
			args: args{
				c: ctx,
			},
			mockFn: func(a args) {
				mockUC.EXPECT().Delete(ctx.Request.Context(), gomock.Any())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockFn(tt.args)
			a := &TodoHandler{
				TodoUsecase: tt.fields.TodoUsecase,
			}
			a.DeleteTodo(tt.args.c)
		})
	}
}
