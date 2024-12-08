package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// 设置一个名为name的cookie，值为Brumble的cookie
		// 设置cookie的过期时间为1 hour
		// 3600是秒，也就是1小时
		c.SetCookie("name", "Brumble", 3600, "/", "", false, true)
		// 获取名为"name"的cookie
		cookieValue, err := c.Cookie("name")
		if err != nil {
			fmt.Println("not found name")
			return
		}
		fmt.Println("name cookie =", cookieValue)
	})
	r.Run(":8080")
}
