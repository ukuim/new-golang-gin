package main

import (
	"fmt"

	"example.com/gin-demo/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var users []models.Users
		models.DB.Find(&users)
		c.JSON(200, gin.H{
			"users": users,
		})
	})

	fmt.Println("127.0.0.1:8080")

	r.Run()
}
