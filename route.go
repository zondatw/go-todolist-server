package main

import "./todos"

func initRoute() {
	router.Use(CORSMiddleware)
	api := router.Group("/api")

	todosService := todos.NewTodosService(db)

	todosRoute := api.Group("/todos")
	{
		todosRoute.GET("/", todosService.GetAll)
		todosRoute.POST("/", todosService.Add)
		todosRoute.GET("/:id", todosService.Get)
		todosRoute.PUT("/:id", todosService.Update)
		todosRoute.DELETE("/:id", todosService.Delete)
	}
}
