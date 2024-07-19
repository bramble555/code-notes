package main

import "github.com/gin-gonic/gin"

type StructA struct {
	FieldA string `form:"field_a"`
}
type StructB struct {
	NestedStrcut StructA
	FieldB       string `form:"field_b"`
}

func main() {
	r := gin.Default()
	// 绑定的是Body里面的form-data
	r.GET("/", func(c *gin.Context) {
		var b StructB
		c.Bind(&b)
		c.JSON(200, gin.H{
			"a": b.NestedStrcut,
			"b": b.FieldB,
		})
		c.JSON(200, "success")
	})
	r.Run(":8080")
}
