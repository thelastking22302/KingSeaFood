package main

import (
	"thelastking/kingseafood/api"
	"thelastking/kingseafood/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/sign-up", api.SignUpHandler(server.Run()))
		users.POST("/sign-in", api.SignInHandler(server.Run()))
	}
	r.Run(":3250")
}
