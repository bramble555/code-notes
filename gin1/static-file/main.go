package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/st", "./static")
	// 在golang总，没有相对文件的路径，它只有相对项目的路径
	// 第一个参数是URL请求这个静态目录的前缀， 第二个参数是一个目录，注意，前缀不要重复
	// url是/static/abc.txt 或者/static/titian.png
	router.StaticFS("/static", http.Dir("./static"))
	// 配置单个文件， 网页请求的路由，文件的路径
	router.StaticFile("/titian.png", "static/titian.png")
	router.Run(":8080")
}
