package main

func initRoute() {
	router.Use(CORSMiddleware)
	api := router.Group("/api")
	todosRoute := api.Group("/todos")
	{
		todosRoute.GET("/", todosGET)
		todosRoute.POST("/", todosPOST)
		todosRoute.GET("/:id", todoGET)
		todosRoute.PUT("/:id", todoPUT)
		todosRoute.DELETE("/:id", todoDELETE)
	}
}
