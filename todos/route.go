package todos

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Route(api *gin.RouterGroup, db *sql.DB) {
	todosService := todosService{db: db}

	todosRoute := api.Group("/todos")
	{
		todosRoute.GET("/", todosService.GetAll)
		todosRoute.POST("/", todosService.Add)
		todosRoute.GET("/:id", todosService.Get)
		todosRoute.PUT("/:id", todosService.Update)
		todosRoute.DELETE("/:id", todosService.Delete)
	}
}
