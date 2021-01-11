package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	allUsers := []user{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(200, allUsers)
	})

	r.Run(":8080")
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}