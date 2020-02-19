package todos

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll(context *gin.Context)
	Add(context *gin.Context)
	Get(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type todosService struct {
	db *sql.DB
}

func (service *todosService) GetAll(context *gin.Context) {
	todos := queryTodoTable(service.db)
	context.JSON(http.StatusOK, todos)
}

func (service *todosService) Add(context *gin.Context) {
	var todo Todo
	var newTodo Todo
	if err := context.ShouldBindJSON(&todo); err == nil {
		newTodo = addTodo(service.db, todo)
		context.JSON(http.StatusOK, newTodo)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (service *todosService) Get(context *gin.Context) {
	id := context.Param("id")
	todo := getTodo(service.db, id)
	context.JSON(http.StatusOK, todo)
}

func (service *todosService) Update(context *gin.Context) {
	id := context.Param("id")
	var todo Todo
	if err := context.ShouldBindJSON(&todo); err == nil {
		newTodo := updateTodo(service.db, id, todo)
		context.JSON(http.StatusOK, newTodo)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (service *todosService) Delete(context *gin.Context) {
	id := context.Param("id")
	deleteTodo(service.db, id)
	context.JSON(http.StatusNoContent, gin.H{})
}

type sortService struct {
	db *sql.DB
}

func (service *sortService) Get(context *gin.Context) {
	sort := getTodoSort(service.db)
	context.JSON(http.StatusOK, sort)
}

func (service *sortService) Update(context *gin.Context) {
	var sort []int
	if err := context.ShouldBindJSON(&sort); err == nil {
		newSort := updateTodoSort(service.db, sort)
		context.JSON(http.StatusOK, newSort)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
