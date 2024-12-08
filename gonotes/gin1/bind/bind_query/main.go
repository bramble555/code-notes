package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	type Person struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}
	r := gin.Default()
	// 绑定的是Params
	r.POST("/", func(c *gin.Context) {
		var person Person
		if c.ShouldBind(&person) == nil {
			log.Println(person.Name)
			log.Println(person.Age)
		}
		c.JSON(200, "success")
	})
	r.Run(":8080")
}
