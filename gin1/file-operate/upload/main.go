package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// 单位是字节， << 是左移预算符号，等价于 8 * 2^20
	// gin对文件上传大小的默认值是32MB
	r.MaxMultipartMemory = 8 << 20 // 8MIB

	// 单文件上传
	// r.POST("/", func(c *gin.Context) {
	// 	file, _ := c.FormFile("file")
	// 	log.Println(file.Filename)
	// 	dst := "./" + file.Filename
	// 	// 文件对象  文件路径，注意要从项目根路径开始写,当然也可以改变工作区
	// 	c.SaveUploadedFile(file, dst)
	// 	c.String(200, "ok")
	// })

	// 多文件上传
	r.POST("/", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			dst := "./file/" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(200, "ok")

	})
	r.Run(":8080")
}
