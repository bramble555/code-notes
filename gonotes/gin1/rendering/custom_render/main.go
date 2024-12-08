package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义函数渲染
func formatAsDate(words string) string {

	return fmt.Sprintf("现在是%s,%s", time.Now().Format("2006-01-02 15:04:05"), words)
}
func main() {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	r.LoadHTMLFiles("./file1.html")
	r.GET("/", func(c *gin.Context) {
		// 在模板中调用函数，传入words参数
		c.HTML(http.StatusOK, "file1.html", gin.H{
			"words": "请珍惜你的时间",
		})
	})
	// 监听的数字前面需要需要加上:
	r.Run(":8080")
}
