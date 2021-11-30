package controllers

//controller atau handler
import (
	"net/http"
	"strconv"

	"github.com/KennyKur/CRUD_Todo/models"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	TodoEntity models.TodoEntity
}

func NewTodoHandler(r *gin.RouterGroup, us models.TodoEntity) {
	handler := &TodoHandler{
		TodoEntity: us,
	}
	r.GET("/Todo/", handler.FindTodos)
	r.GET("/Todo/:id", handler.FindTodo)
	r.POST("/Todos", handler.CreateTodo)
	r.PATCH("Todo/update/:id", handler.UpdateTodo)
	r.DELETE("Todo/delete/:id", handler.DeleteTodo)
}
func (a *TodoHandler) FindTodos(c *gin.Context) {
	todos, _ := a.TodoEntity.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": todos})
}

func (a *TodoHandler) FindTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := a.TodoEntity.GetByID(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": todo})
}

func (a *TodoHandler) CreateTodo(c *gin.Context) {
	var input models.User_todo_list
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.TodoEntity.Create(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "data berhasil ditambahkan"})
}

func (a *TodoHandler) UpdateTodo(c *gin.Context) {
	var input models.User_todo_list
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.TodoEntity.Update(c.Request.Context(), input, int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "data berhasil diubah"})

}

func (a *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.TodoEntity.Delete(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "data berhasil dihapus"})

}
