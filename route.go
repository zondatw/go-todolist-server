package main

import (
	"github.com/zondaTW/go-todolist-server/middleware"
	"github.com/zondaTW/go-todolist-server/todos"
	"github.com/zondaTW/go-todolist-server/user"
)

func initRoute(authKey string) {
	router.Use(middleware.CORSMiddleware)

	jwtMiddleware := middleware.GetJWTMiddleware(authKey, user.GetJWTAuthFunc(db))
	userRouter := router.Group("/user")
	userRouter.POST("/login", jwtMiddleware.LoginHandler)
	userRouter.POST("/register", user.AddUser(db))

	auth := router.Group("/auth")
	auth.Use(jwtMiddleware.MiddlewareFunc())
	auth.GET("/refresh_token", jwtMiddleware.RefreshHandler)

	api := router.Group("/api")
	api.Use(jwtMiddleware.MiddlewareFunc())
	todos.Route(api, db)
}
