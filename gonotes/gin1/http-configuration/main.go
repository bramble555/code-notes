package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义http配置
func main() {
	// 配置log颜色
	gin.ForceConsoleColor()
	r := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "ok",
		})
	})
	s.ListenAndServe()
}
