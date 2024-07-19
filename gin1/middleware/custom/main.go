package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	// 这里的代码是程序一开始就会执行
	return func(c *gin.Context) {
		// 这里是请求来了才会执行
		t := time.Now()
		// 中间件传递数据，后面middleware如果需要
		// 设置name可变参数
		c.Set("name", "Bramble")
		// 下一个middleware
		c.Next()
		// 中断下一个middleware
		// c.Abort()
		// 计算middleware用了多久
		latencyTime := time.Since(t)
		log.Println(latencyTime)
		// 目前发送的状态
		status := c.Writer.Status()
		log.Println(status) // 200
	}
}
func main() {
	// 自定义log输出
	r := gin.New()
	// global middleware
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s]===%s  %s  状态码%d  %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.StatusCode,
			param.ErrorMessage,
		)
	},
	))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(Logger())
	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
		fmt.Println(c.Get("name"))
	})
	r.Run(":8080")

}
