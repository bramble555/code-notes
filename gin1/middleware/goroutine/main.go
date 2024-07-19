package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		cC := c.Copy()
		go func() {
			time.Sleep(1 * time.Second)
		}()
		// 请注意您使用的是复制的上下文 "cC"，这一点很重要
		log.Println("Done! in path "+cC.Request.URL.Path, "我是副本！！")
	})
	r.GET("/2", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(1 * time.Second)

		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path "+c.Request.URL.Path, "我是正品")
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
