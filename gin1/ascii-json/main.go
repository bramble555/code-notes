package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化引擎
	r := gin.Default()
	// 定义路由
	r.GET("/", func(c *gin.Context) {
		data := map[string]interface{}{
			"name": "ccl",
			"age":  18,
		}
		// c.JSON(200, data) // Content-Type = application/json; charset=utf-8
		// unicode to ASCII
		// sets the Content-Type as "application/json"
		c.AsciiJSON(http.StatusOK, data) //Content-Type = application/json
	})
	r.Run(":8080")
}
