package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 渲染模板
func main() {
	r := gin.Default()
	// 传输一个或者几个文件
	// r.LoadHTMLFiles("templates/template1.html")
	// 传输一个目录下的文件
	// 传输多个路由
	r.LoadHTMLGlob("templates/*")
	r.GET("/template1.html", func(c *gin.Context) {
		// 注意HTML方法传fileName的时候，取代的是上面LoadHTMLGlob方法里面的*
		c.HTML(http.StatusOK, "template1.html", gin.H{
			"title": "Main website1",
		})
	})
	r.GET("/template2.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template2.html", gin.H{
			"title": "Main website2",
		})
	})
	r.Run(":8080")
}
