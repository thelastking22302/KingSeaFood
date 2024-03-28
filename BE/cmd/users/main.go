package main

import (
	"thelastking/kingseafood/api"
	"thelastking/kingseafood/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	server.Run()

	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/sign-up", api.SignUpHandler(&gorm.DB{}))
	}
	r.Run(":3250")
}
