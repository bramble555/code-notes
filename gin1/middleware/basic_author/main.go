package main

import (
	"github.com/gin-gonic/gin"
)

// 前后端最流行的jwt为例
// 如果用户登录了，前端发来的每一次请求都会在请求头上携带上token
// 后台拿到这个token进行校验，验证是否过期，是否非法
// 如果通过就说明这个用户是登录过的
// 不通过就说明用户没有登录
func basicAuthor(c *gin.Context) {
	tokenName := c.GetHeader("token")
	if tokenName == "Bramble" {
		// 验证通过
		c.Next()
		return
	}
	c.JSON(200, gin.H{
		"msg": "验证不通过",
	})
	c.Abort()
}
func main() {
	r := gin.Default()
	apiUser := r.Group("")
	{
		apiUser.POST("/login", func(c *gin.Context) {
			c.String(200, "注册成功")
		})
	}
	apiHome := r.Group("/entry").Use(basicAuthor)
	{
		apiHome.POST("/home", func(c *gin.Context) {
			c.String(200, "登录成功,欢迎回来")
		})
	}
	r.Run(":8080")
}
