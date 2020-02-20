package main

import "github.com/zondaTW/go-todolist-server/todos"

func initRoute() {
	router.Use(CORSMiddleware)
	api := router.Group("/api")
	todos.Route(api, db)
}
