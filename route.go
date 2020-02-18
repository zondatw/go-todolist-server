package main

import "./todos"

func initRoute() {
	router.Use(CORSMiddleware)
	api := router.Group("/api")
	todos.Route(api, db)
}
