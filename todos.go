package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func todosGET(context *gin.Context) {
	todos := queryTodoTable(db)
	context.JSON(http.StatusOK, todos)
}

func todosPOST(context *gin.Context) {
	var todo Todo
	var newTodo Todo
	if err := context.ShouldBindJSON(&todo); err == nil {
		newTodo = addTodo(db, todo)
		context.JSON(http.StatusOK, newTodo)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func todoGET(context *gin.Context) {
	id := context.Param("id")
	todo := getTodo(db, id)
	context.JSON(http.StatusOK, todo)
}

func todoPUT(context *gin.Context) {
	id := context.Param("id")
	var todo Todo
	if err := context.ShouldBindJSON(&todo); err == nil {
		newTodo := updateTodo(db, id, todo)
		context.JSON(http.StatusOK, newTodo)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func todoDELETE(context *gin.Context) {
	id := context.Param("id")
	deleteTodo(db, id)
	context.JSON(http.StatusNoContent, gin.H{})
}
