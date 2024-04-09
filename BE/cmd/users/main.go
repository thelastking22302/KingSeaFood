package main

import (
	"thelastking/kingseafood/api"
	"thelastking/kingseafood/middleware"
	"thelastking/kingseafood/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", api.SignUpHandler(server.Run()))
		auth.POST("/sign-in", api.SignInHandler(server.Run()))
	}
	users := r.Group("/users", middleware.RefreshJWTMiddleware(), middleware.JWTMiddleware())
	{
		users.GET("/profile/id", api.ProfileUser(server.Run()))
		users.PATCH("/update/id", api.UpdateUserHandler(server.Run()))
		users.DELETE("/delete/id", api.DeletedUserHandler(server.Run()))
	}
	r.Run(":3250")
}
