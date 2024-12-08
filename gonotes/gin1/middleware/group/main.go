package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 分组，方便管理,也方便使用全局middleware
	// 当然也可以在aGroup进行nested group(嵌套分组)
	aGroup := r.Group("/a")
	aGroup.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})
	r.Run(":8080")

}
