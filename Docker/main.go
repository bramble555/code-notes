package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!") // 返回字符串内容
	})
	r.Run(":8888") // 启动服务器，监听8888端口
}
