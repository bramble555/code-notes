package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	type User struct {
		// 注意tag不要写错
		// binnding是内置规则，需要校验是否满足条件
		Name string `json:"name" uri:"name" binding:"required,min=3"`
		Age  int    `json:"age" uri:"age" binding:"required,gt=18" msg:"年龄需要大于等于18岁"`
	}
	r := gin.Default()
	// json格式
	r.POST("/json", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(505, gin.H{
				"msg": "出错了",
			})
		} else {
			c.JSON(200, gin.H{
				"name": user.Name,
				"age":  user.Age,
			})
		}

	})
	// 对应tag里面的uri
	r.GET("/uri/:name/:age", func(c *gin.Context) {
		var user User
		err := c.ShouldBindUri(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"msg": "你错了",
			})
			return
		} else {
			fmt.Println(user)
			c.JSON(200, gin.H{
				"msg": "ok",
			})
		}

	})
	r.Run(":8080")
}
